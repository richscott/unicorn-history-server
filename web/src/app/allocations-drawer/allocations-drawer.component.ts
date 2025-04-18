import {
  Component,
  EventEmitter,
  Injectable,
  Input,
  OnInit,
  Output,
  ViewChild,
} from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatDrawer } from '@angular/material/sidenav';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { AllocationInfo } from '@app/models/alloc-info.model';
import { AppInfo } from '@app/models/app-info.model';
import { ColumnDef } from '@app/models/column-def.model';
import { EnvConfigService } from '@app/services/envconfig/envconfig.service';
import { CommonUtil } from '@app/utils/common.util';

@Injectable()
@Component({
  selector: 'app-allocations-drawer',
  templateUrl: './allocations-drawer.component.html',
  styleUrls: ['./allocations-drawer.component.scss'],
})
export class AllocationsDrawerComponent implements OnInit {
  @ViewChild('matDrawer', { static: false }) matDrawer!: MatDrawer;
  @ViewChild('allocationMatPaginator', { static: true }) allocPaginator!: MatPaginator;
  @ViewChild('allocSort', { static: true }) allocSort!: MatSort;
  @Input() allocDataSource!: MatTableDataSource<AllocationInfo & { expanded: boolean }>;
  @Input() selectedRow!: AppInfo | null;
  @Input() externalLogsBaseUrl!: string | null;
  @Input() partitionSelected!: string;
  @Input() leafQueueSelected!: string;

  @Output() removeRowSelection = new EventEmitter<void>();

  allocColumnDef: ColumnDef[] = [];
  allocColumnIds: string[] = [];
  selectedAllocationsRow: number = -1;

  ngOnChanges(): void {
    if (this.allocDataSource) {
      this.allocDataSource.paginator = this.allocPaginator;
      this.allocDataSource.sort = this.allocSort;
    }
  }
  constructor(private envConfig: EnvConfigService) {}

  ngOnInit(): void {
    this.allocColumnDef = [
      { colId: 'displayName', colName: 'Display Name', colWidth: 1 },
      { colId: 'allocationKey', colName: 'Allocation Key', colWidth: 1 },
      { colId: 'nodeId', colName: 'Node ID', colWidth: 1 },
      {
        colId: 'log',
        colName: 'Log Link',
        colWidth: 1,
      },
      {
        colId: 'resource',
        colName: 'Resource',
        colFormatter: CommonUtil.resourceColumnFormatter,
        colWidth: 1,
      },
      { colId: 'priority', colName: 'Priority', colWidth: 0.5 },
      {
        colId: 'elapsedTime',
        colName: 'Time to Allocation',
        colWidth: 0.8,
      },
    ];
    this.allocColumnIds = this.allocColumnDef.map((col) => col.colId);
    this.externalLogsBaseUrl = this.envConfig.getExternalLogsBaseUrl();
  }

  formatResources(colValue: string): string[] {
    const arr: string[] = colValue.split('<br/>');
    // Check if there are "cpu" or "Memory" elements in the array
    const hasCpu = arr.some((item) => item.toLowerCase().includes('cpu'));
    const hasMemory = arr.some((item) => item.toLowerCase().includes('memory'));
    if (!hasCpu) {
      arr.unshift('CPU: n/a');
    }
    if (!hasMemory) {
      arr.unshift('Memory: n/a');
    }

    // Concatenate the two arrays, with "cpu" and "Memory" elements first
    const cpuAndMemoryElements = arr.filter(
      (item) => item.toLowerCase().includes('CPU') || item.toLowerCase().includes('Memory')
    );
    const otherElements = arr.filter(
      (item) => !item.toLowerCase().includes('CPU') && !item.toLowerCase().includes('Memory')
    );
    const result = cpuAndMemoryElements.concat(otherElements);

    return result;
  }

  isAllocDataSourceEmpty() {
    return this.allocDataSource?.data && this.allocDataSource.data.length === 0;
  }

  allocationsDetailToggle(row: number) {
    if (this.selectedAllocationsRow !== -1) {
      if (this.selectedAllocationsRow !== row) {
        this.allocDataSource.data[this.selectedAllocationsRow].expanded = false;
        this.selectedAllocationsRow = row;
        this.allocDataSource.data[row].expanded = true;
      } else {
        this.allocDataSource.data[this.selectedAllocationsRow].expanded = false;
        this.selectedAllocationsRow = -1;
      }
    } else {
      this.selectedAllocationsRow = row;
      this.allocDataSource.data[row].expanded = true;
    }
  }

  closeDrawer() {
    this.selectedAllocationsRow = -1;
    this.matDrawer.close();
    this.removeRowSelection.emit();
  }

  openDrawer() {
    this.matDrawer.open();
  }

  logClick(e: MouseEvent) {
    e.stopPropagation();
  }

  copyLinkToClipboard() {
    const url = window.location.href.split('?')[0];
    const copyString = `${url}?partition=${this.partitionSelected}&queue=${this.leafQueueSelected}&applicationId=${this?.selectedRow?.applicationId}`;
    navigator.clipboard
      .writeText(copyString)
      .catch((error) => console.error('Writing to the clipboard is not allowed. ', error));
  }

  calculateElapsedTime(requestTime: string, allocationTime: string): string {
    if (!requestTime) {
      return 'n/a';
    }

    const request = parseInt(requestTime) / 1_000_000; // nanoseconds to milliseconds
    const allocation = parseInt(allocationTime) / 1_000_000; // nanoseconds to milliseconds
    const diffInSeconds = Math.floor((allocation - request) / 1000);

    return `${diffInSeconds}s`;
  }
}

import { Component, OnInit } from "@angular/core";
import { CommonModule } from "@angular/common";
import { ApiService } from "../../services/api.service";

@Component({
  selector: "app-report-viewer",
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="report-viewer">
      <h2>Report Viewer</h2>
      <div class="report-list" *ngIf="reports.length > 0">
        <div *ngFor="let report of reports" class="report-card">
          <h3>{{ report.type | titlecase }} Report</h3>
          <p>Status: {{ report.status }}</p>
          <button (click)="downloadReport(report.id)">Download</button>
        </div>
      </div>
      <div *ngIf="reports.length === 0"><p>No reports available</p></div>
    </div>
  `,
  styles: [`
    .report-viewer { padding: 1rem; }
    .report-card { background: white; padding: 1.5rem; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); margin-bottom: 1rem; }
    button { padding: 0.5rem 1rem; background: #1a1a2e; color: white; border: none; border-radius: 4px; cursor: pointer; }
  `]
})
export class ReportViewerComponent implements OnInit {
  reports: any[] = [];
  constructor(private apiService: ApiService) {}
  ngOnInit(): void { this.loadReports(); }
  loadReports(): void {
    this.apiService.getReports().subscribe({ next: (data: any) => { this.reports = data; }, error: (err: any) => { console.error(err); } });
  }
  downloadReport(id: string): void {
    this.apiService.downloadReport(id).subscribe({ next: () => { console.log("Download started"); }, error: (err: any) => { console.error(err); } });
  }
}

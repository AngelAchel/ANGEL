import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService } from '../../services/api.service';

@Component({
  selector: 'app-console',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="console">
      <h2>Agent Console</h2>
      <div class="console-grid">
        <div class="console-panel">
          <h3>Active Agents</h3>
          <ul>
            <li *ngFor="let agent of agents">{{ agent.hostname || agent.id }}</li>
          </ul>
        </div>
        <div class="console-panel">
          <h3>Tasks</h3>
          <ul>
            <li *ngFor="let task of tasks">{{ task.id }} - {{ task.type }}</li>
          </ul>
        </div>
        <div class="console-panel">
          <h3>Results</h3>
          <ul>
            <li *ngFor="let result of results">{{ result.id }} - {{ result.success ? '✅' : '❌' }}</li>
          </ul>
        </div>
      </div>
    </div>
  `,
  styles: [`
    .console { padding: 20px; }
    .console-grid { display: grid; grid-template-columns: 1fr 1fr 1fr; gap: 20px; }
    .console-panel { background: #1e1e1e; color: #00ff00; padding: 15px; border-radius: 8px; }
    h3 { color: #00ff88; margin-top: 0; }
    ul { list-style: none; padding: 0; }
    li { padding: 4px 0; font-family: monospace; }
  `]
})
export class AgentConsoleComponent implements OnInit {
  agents: any[] = [];
  tasks: any[] = [];
  results: any[] = [];

  constructor(private api: ApiService) {}

  ngOnInit(): void {
    this.api.getAgents().subscribe({ next: r => this.agents = r, error: () => {} });
    this.api.getTasks().subscribe({ next: r => this.tasks = r, error: () => {} });
    this.api.getResults().subscribe({ next: r => this.results = r, error: () => {} });
  }
}

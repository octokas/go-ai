# Mixed Menu Tables

```html
<div class="table-advanced-layout">
  <!-- Context Bar -->
  <div class="context-bar">
    <div class="breadcrumb">
      <span>Projects</span>
      <svg><!-- Chevron --></svg>
      <span>Active</span>
    </div>

    <div class="context-actions">
      <button class="btn-secondary">Export</button>
      <button class="btn-primary">New Project</button>
    </div>
  </div>

  <div class="table-container">
    <!-- Side Tools -->
    <div class="table-tools">
      <div class="tool-section">
        <button class="tool-button" data-tooltip="Filter">
          <svg><!-- Filter icon --></svg>
        </button>
        <button class="tool-button" data-tooltip="Group">
          <svg><!-- Group icon --></svg>
        </button>
        <button class="tool-button" data-tooltip="Sort">
          <svg><!-- Sort icon --></svg>
        </button>
      </div>

      <div class="tool-section">
        <button class="tool-button" data-tooltip="Settings">
          <svg><!-- Settings icon --></svg>
        </button>
      </div>
    </div>

    <!-- Main Table Area -->
    <div class="table-content">
      <table class="table">
        <!-- Table content -->
      </table>
    </div>

    <!-- Detail Panel -->
    <div class="detail-panel">
      <!-- Selected row details -->
    </div>
  </div>
</div>
```

```scss
.table-advanced-layout {
  --tools-width: 48px;
  --detail-width: 320px;

  height: 100%;
  display: flex;
  flex-direction: column;

  .context-bar {
    height: 48px;
    padding: var(--space-4);
    display: flex;
    align-items: center;
    justify-content: space-between;
    background: var(--color-surface);
    border-bottom: 1px solid var(--color-border);
  }

  .table-container {
    flex: 1;
    display: flex;
    overflow: hidden;
  }

  .table-tools {
    width: var(--tools-width);
    border-right: 1px solid var(--color-border);
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding: var(--space-2);

    .tool-section {
      display: flex;
      flex-direction: column;
      gap: var(--space-2);
    }

    .tool-button {
      width: 32px;
      height: 32px;
      border-radius: var(--radius-md);

      &:hover {
        background: var(--color-surface-hover);
      }

      &.active {
        background: var(--color-primary);
        color: white;
      }
    }
  }

  .detail-panel {
    width: var(--detail-width);
    border-left: 1px solid var(--color-border);
    background: var(--color-surface);
    transform: translateX(100%);
    transition: transform var(--duration-normal) var(--ease-spring);

    &.open {
      transform: translateX(0);
    }
  }
}

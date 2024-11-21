# Layered Tables

```html
<div class="table-layered">
  <!-- Master Table -->
  <div class="table-layer master">
    <div class="layer-header">
      <h3>Projects</h3>
      <div class="layer-actions">
        <button class="btn-secondary">Filter</button>
        <button class="btn-primary">Add Project</button>
      </div>
    </div>

    <table class="table">
      <!-- Master table content -->
    </table>
  </div>

  <!-- Detail Table -->
  <div class="table-layer detail">
    <div class="layer-header">
      <div class="layer-navigation">
        <button class="btn-icon back">
          <svg><!-- Back icon --></svg>
        </button>
        <h3>Project Tasks</h3>
      </div>

      <div class="layer-actions">
        <button class="btn-primary">Add Task</button>
      </div>
    </div>

    <table class="table">
      <!-- Detail table content -->
    </table>
  </div>

  <!-- Sub-detail Layer -->
  <div class="table-layer sub-detail">
    <div class="layer-header">
      <div class="layer-navigation">
        <button class="btn-icon back">
          <svg><!-- Back icon --></svg>
        </button>
        <h3>Task Details</h3>
      </div>
    </div>

    <div class="detail-content">
      <!-- Detail content -->
    </div>
  </div>
</div>
```

```scss
.table-layered {
  --layer-width: 100%;
  --layer-transition: var(--duration-normal) var(--ease-spring);

  position: relative;
  height: 100%;
  overflow: hidden;

  .table-layer {
    position: absolute;
    inset: 0;
    background: var(--color-surface);
    transition: transform var(--layer-transition);

    &.detail {
      transform: translateX(100%);

      &.active {
        transform: translateX(0);
      }
    }

    &.sub-detail {
      transform: translateX(200%);

      &.active {
        transform: translateX(0);
      }
    }
  }

  .layer-header {
    height: 64px;
    padding: var(--space-4);
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: 1px solid var(--color-border);
  }

  .layer-navigation {
    display: flex;
    align-items: center;
    gap: var(--space-3);
  }

  // Animation states
  &[data-active-layer="detail"] {
    .master {
      transform: translateX(-30%);
      opacity: 0.5;
    }
  }

  &[data-active-layer="sub-detail"] {
    .master {
      transform: translateX(-60%);
      opacity: 0.3;
    }
    .detail {
      transform: translateX(-30%);
      opacity: 0.5;
    }
  }
}

# Top/Bottom Menu Tables

```html
<div class="table-with-menus">
  <!-- Top Menu -->
  <div class="table-top-menu">
    <div class="menu-section">
      <div class="bulk-actions">
        <select class="bulk-select">
          <option>Bulk Actions</option>
          <option>Delete Selected</option>
          <option>Change Status</option>
        </select>
        <button class="btn-secondary">Apply</button>
      </div>

      <div class="view-options">
        <div class="btn-group">
          <button class="btn-icon active">
            <svg><!-- Grid view icon --></svg>
          </button>
          <button class="btn-icon">
            <svg><!-- List view icon --></svg>
          </button>
        </div>
      </div>
    </div>

    <div class="menu-section">
      <div class="table-search">
        <input type="search" placeholder="Search...">
      </div>

      <button class="btn-primary">Add New</button>
    </div>
  </div>

  <!-- Table Content -->
  <div class="table-wrapper">
    <table class="table">
      <!-- Table content -->
    </table>
  </div>

  <!-- Bottom Menu -->
  <div class="table-bottom-menu">
    <div class="table-info">
      <span>Selected: 3 of 50 items</span>
    </div>

    <div class="table-pagination">
      <!-- Pagination controls -->
    </div>
  </div>
</div>
```

```scss
.table-with-menus {
  --menu-height: 64px;

  display: flex;
  flex-direction: column;
  height: 100%;

  .table-top-menu,
  .table-bottom-menu {
    height: var(--menu-height);
    padding: var(--space-4);
    background: var(--color-surface);
    border-bottom: 1px solid var(--color-border);

    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: var(--space-4);
  }

  .table-wrapper {
    flex: 1;
    overflow: auto;
  }

  .menu-section {
    display: flex;
    align-items: center;
    gap: var(--space-4);
  }
}

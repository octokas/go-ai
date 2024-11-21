# Side Menu Tables

```html
<div class="table-with-side-menu">
  <aside class="table-side-menu">
    <div class="side-menu-header">
      <h3>Filters</h3>
      <button class="btn-icon collapse-menu">
        <svg><!-- Collapse icon --></svg>
      </button>
    </div>

    <div class="side-menu-content">
      <div class="filter-group">
        <label class="filter-label">Categories</label>
        <div class="filter-options">
          <label class="checkbox-label">
            <input type="checkbox" checked>
            <span>Active (23)</span>
          </label>
          <!-- More filter options -->
        </div>
      </div>

      <div class="filter-group">
        <label class="filter-label">Date Range</label>
        <div class="date-range-picker">
          <!-- Date picker implementation -->
        </div>
      </div>
    </div>

    <div class="side-menu-footer">
      <button class="btn-secondary">Reset</button>
      <button class="btn-primary">Apply Filters</button>
    </div>
  </aside>

  <div class="table-main-content">
    <table class="table">
      <!-- Table content -->
    </table>
  </div>
</div>
```

```scss
.table-with-side-menu {
  --side-menu-width: 280px;
  --side-menu-collapsed: 60px;

  display: flex;
  gap: var(--space-4);
  height: 100%;

  .table-side-menu {
    width: var(--side-menu-width);
    border-right: 1px solid var(--color-border);
    transition: width var(--duration-normal) var(--ease-spring);

    &.collapsed {
      width: var(--side-menu-collapsed);

      .filter-group {
        opacity: 0;
        pointer-events: none;
      }
    }
  }

  .table-main-content {
    flex: 1;
    min-width: 0; // Prevent flex item overflow
  }
}

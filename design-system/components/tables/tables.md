# Tables

## Basic Table
```html
<div class="table-container">
  <table class="table">
    <thead>
      <tr>
        <th class="table-header sortable">
          Name
          <svg class="sort-icon"><!-- Sort icon SVG --></svg>
        </th>
        <th class="table-header">Status</th>
        <th class="table-header text-right">Actions</th>
      </tr>
    </thead>
    <tbody>
      <tr class="table-row">
        <td class="table-cell">
          <div class="cell-content">
            <img src="avatar.jpg" class="cell-avatar" />
            <span class="cell-text">John Doe</span>
          </div>
        </td>
        <td class="table-cell">
          <span class="status-badge success">Active</span>
        </td>
        <td class="table-cell text-right">
          <div class="action-buttons">
            <button class="btn-icon">Edit</button>
            <button class="btn-icon">Delete</button>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

```scss
.table-container {
  --table-border: var(--color-border);
  --header-bg: var(--color-surface-alt);
  --row-hover: var(--color-surface-hover);

  width: 100%;
  overflow-x: auto;
  border: 1px solid var(--table-border);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-sm);
}

.table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0;

  .table-header {
    padding: var(--space-4);
    background: var(--header-bg);
    font-weight: var(--font-semibold);
    text-align: left;
    white-space: nowrap;

    &.sortable {
      cursor: pointer;
      user-select: none;

      &:hover {
        background: var(--color-surface-hover);
      }

      .sort-icon {
        transition: transform var(--duration-fast) var(--ease-spring);
      }

      &.sorted-asc .sort-icon {
        transform: rotate(180deg);
      }
    }
  }

  .table-row {
    transition: background var(--duration-fast) var(--ease-default);

    &:hover {
      background: var(--row-hover);
    }

    &:not(:last-child) {
      border-bottom: 1px solid var(--table-border);
    }
  }

  .table-cell {
    padding: var(--space-4);
    vertical-align: middle;
  }
}

// Advanced Table Features
.table-advanced {
  // Sticky Header
  thead {
    position: sticky;
    top: 0;
    z-index: 10;
    background: var(--header-bg);
  }

  // Row Selection
  .row-selector {
    width: 48px;
    text-align: center;
  }

  // Expandable Rows
  .row-expander {
    cursor: pointer;

    &.expanded {
      background: var(--color-surface-selected);
    }
  }

  // Nested Tables
  .nested-table {
    margin: var(--space-4);
    background: var(--color-surface-alt);
    border-radius: var(--radius-md);
  }
}
```

## Interactive Table Features
```html
<div class="table-toolbar">
  <div class="table-search">
    <input type="search" placeholder="Search..." />
  </div>

  <div class="table-actions">
    <button class="btn-secondary">Export</button>
    <button class="btn-primary">Add New</button>
  </div>
</div>

<div class="table-footer">
  <div class="table-info">
    Showing 1-10 of 100 items
  </div>

  <div class="table-pagination">
    <button class="pagination-btn" disabled>Previous</button>
    <div class="pagination-numbers">
      <button class="pagination-number active">1</button>
      <button class="pagination-number">2</button>
      <span class="pagination-ellipsis">...</span>
      <button class="pagination-number">10</button>
    </div>
    <button class="pagination-btn">Next</button>
  </div>
</div>
```

```scss
.table-toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--space-4);
  border-bottom: 1px solid var(--table-border);
}

.table-pagination {
  display: flex;
  align-items: center;
  gap: var(--space-2);

  .pagination-number {
    min-width: 32px;
    height: 32px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: var(--radius-md);

    &.active {
      background: var(--color-primary);
      color: white;
    }

    &:hover:not(.active) {
      background: var(--color-surface-hover);
    }
  }
}

// Responsive Table
@media (max-width: 768px) {
  .table-responsive {
    .table-header {
      display: none;
    }

    .table-row {
      display: block;
      padding: var(--space-4);

      &:not(:last-child) {
        border-bottom: 1px solid var(--table-border);
      }
    }

    .table-cell {
      display: flex;
      justify-content: space-between;
      padding: var(--space-2) 0;

      &::before {
        content: attr(data-label);
        font-weight: var(--font-semibold);
      }
    }
  }
}
```

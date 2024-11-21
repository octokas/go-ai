# Editable Tables

## Inline Editable Table
```html
<table class="table table-editable">
  <thead>
    <tr>
      <th>Name</th>
      <th>Email</th>
      <th>Role</th>
      <th>Status</th>
      <th>Actions</th>
    </tr>
  </thead>
  <tbody>
    <tr class="editable-row" data-row-id="1">
      <td class="editable-cell" data-field="name">
        <div class="cell-view">John Doe</div>
        <div class="cell-edit">
          <input type="text" class="edit-input" value="John Doe">
        </div>
      </td>
      <td class="editable-cell" data-field="email">
        <div class="cell-view">john@example.com</div>
        <div class="cell-edit">
          <input type="email" class="edit-input" value="john@example.com">
        </div>
      </td>
      <td class="editable-cell" data-field="role">
        <div class="cell-view">Admin</div>
        <div class="cell-edit">
          <select class="edit-select">
            <option>Admin</option>
            <option>User</option>
            <option>Guest</option>
          </select>
        </div>
      </td>
      <td class="editable-cell" data-field="status">
        <div class="cell-view">
          <span class="status-badge active">Active</span>
        </div>
        <div class="cell-edit">
          <div class="toggle-switch">
            <input type="checkbox" checked>
            <span class="toggle-slider"></span>
          </div>
        </div>
      </td>
      <td class="action-cell">
        <div class="edit-actions">
          <button class="btn-icon edit-start">
            <svg><!-- Edit icon --></svg>
          </button>
          <div class="edit-controls">
            <button class="btn-icon save">
              <svg><!-- Save icon --></svg>
            </button>
            <button class="btn-icon cancel">
              <svg><!-- Cancel icon --></svg>
            </button>
          </div>
        </div>
      </td>
    </tr>
  </tbody>
</table>
```

```scss
.table-editable {
  --edit-transition: var(--duration-normal) var(--ease-spring);

  .editable-cell {
    position: relative;
    padding: 0;

    .cell-view,
    .cell-edit {
      padding: var(--space-4);
      transition: opacity var(--edit-transition),
                  transform var(--edit-transition);
    }

    .cell-edit {
      position: absolute;
      inset: 0;
      opacity: 0;
      transform: translateY(-4px);
      pointer-events: none;
    }

    &.editing {
      .cell-view {
        opacity: 0;
        transform: translateY(4px);
        pointer-events: none;
      }

      .cell-edit {
        opacity: 1;
        transform: translateY(0);
        pointer-events: auto;
      }
    }
  }

  .edit-input,
  .edit-select {
    width: 100%;
    padding: var(--space-2);
    border: 1px solid var(--color-border);
    border-radius: var(--radius-md);
    background: var(--color-surface);

    &:focus {
      outline: none;
      border-color: var(--color-primary);
      box-shadow: 0 0 0 2px var(--color-primary-light);
    }
  }

  .edit-actions {
    display: flex;
    gap: var(--space-2);

    .edit-controls {
      display: none;
    }

    &.editing {
      .edit-start { display: none; }
      .edit-controls { display: flex; }
    }
  }
}

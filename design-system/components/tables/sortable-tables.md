# Sortable Tables

```html
<table class="table table-sortable">
  <thead>
    <tr>
      <th class="sortable-header" data-sort="name">
        <div class="header-content">
          Name
          <div class="sort-controls">
            <svg class="sort-arrow up"><!-- Up arrow --></svg>
            <svg class="sort-arrow down"><!-- Down arrow --></svg>
          </div>
        </div>
      </th>
      <th class="sortable-header" data-sort="date">
        <div class="header-content">
          Date
          <div class="sort-controls">
            <span class="sort-index">2</span>
          </div>
        </div>
      </th>
      <!-- More headers -->
    </tr>
  </thead>
  <tbody class="sortable-body">
    <!-- Table rows -->
  </tbody>
</table>
```

```scss
.table-sortable {
  --sort-arrow-size: 12px;
  --sort-transition: var(--duration-fast) var(--ease-spring);

  .sortable-header {
    cursor: pointer;
    user-select: none;

    .header-content {
      display: flex;
      align-items: center;
      gap: var(--space-2);
    }

    .sort-controls {
      display: flex;
      flex-direction: column;
      height: var(--sort-arrow-size);
      opacity: 0.3;
      transition: opacity var(--sort-transition);
    }

    .sort-arrow {
      width: var(--sort-arrow-size);
      height: calc(var(--sort-arrow-size) / 2);
      transition: transform var(--sort-transition);
    }

    .sort-index {
      font-size: var(--text-xs);
      font-weight: var(--font-bold);
      color: var(--color-primary);
    }

    &:hover .sort-controls {
      opacity: 0.7;
    }

    &.sorted {
      .sort-controls { opacity: 1; }
    }

    &.sorted-asc .sort-arrow.up,
    &.sorted-desc .sort-arrow.down {
      color: var(--color-primary);
      transform: scale(1.2);
    }
  }

  // Multi-sort support
  &.multi-sort {
    .sortable-header.sorted {
      background: var(--color-surface-selected);
    }
  }
}

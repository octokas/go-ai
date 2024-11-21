# Menus

## Bento Menu
```html
<div class="bento-menu">
  <div class="bento-grid">
    <button class="bento-item">
      <svg class="bento-icon"><!-- Icon SVG --></svg>
      <span class="bento-label">Item 1</span>
    </button>
    <!-- Repeat for other items -->
  </div>
</div>
```

```scss
.bento-menu {
  --bento-gap: var(--space-2);
  --bento-item-size: 80px;

  .bento-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: var(--bento-gap);
  }

  .bento-item {
    aspect-ratio: 1;
    padding: var(--space-3);
    border-radius: var(--radius-lg);
    background: var(--color-surface);
    transition: all var(--duration-fast) var(--ease-spring);

    &:hover {
      background: var(--color-surface-hover);
      transform: scale(1.02);
    }
  }
}
```

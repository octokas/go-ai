## Döner Menu
```html
<div class="doner-menu">
  <button class="doner-trigger">
    <span class="doner-icon"></span>
  </button>
  <div class="doner-content">
    <div class="doner-item">Item 1</div>
    <!-- Repeat for other items -->
  </div>
</div>
```

```scss
.doner-menu {
  position: relative;

  .doner-content {
    position: absolute;
    top: 100%;
    right: 0;
    min-width: 200px;
    padding: var(--space-2);
    background: var(--color-surface);
    border-radius: var(--radius-lg);
    box-shadow: var(--shadow-lg);
    transform-origin: top right;
    animation: doner-enter var(--duration-normal) var(--ease-spring);
  }
}

@keyframes doner-enter {
  from {
    opacity: 0;
    transform: scale(0.95);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}

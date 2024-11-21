# Drawer Menu

```html
<div class="drawer-menu" data-position="left">
  <div class="drawer-overlay"></div>
  <div class="drawer-content">
    <header class="drawer-header">
      <h2 class="drawer-title">Menu</h2>
      <button class="drawer-close">
        <svg><!-- Close icon SVG --></svg>
      </button>
    </header>

    <nav class="drawer-navigation">
      <ul class="drawer-list">
        <li class="drawer-item">
          <a href="#" class="drawer-link">
            <svg class="drawer-icon"><!-- Icon SVG --></svg>
            <span>Menu Item</span>
          </a>
        </li>
        <!-- More items -->
      </ul>
    </nav>

    <footer class="drawer-footer">
      <!-- Footer content -->
    </footer>
  </div>
</div>
```

```scss
.drawer-menu {
  --drawer-width: 320px;
  --drawer-background: var(--color-surface);

  position: fixed;
  inset: 0;
  z-index: var(--z-drawer);
  visibility: hidden;

  &[data-position="left"] .drawer-content {
    left: 0;
    transform: translateX(-100%);
  }

  &[data-position="right"] .drawer-content {
    right: 0;
    transform: translateX(100%);
  }

  &.active {
    visibility: visible;

    .drawer-overlay {
      opacity: 1;
    }

    .drawer-content {
      transform: translateX(0);
    }
  }
}

.drawer-overlay {
  position: absolute;
  inset: 0;
  background: rgba(0, 0, 0, 0.4);
  opacity: 0;
  transition: opacity var(--duration-normal) var(--ease-default);
}

.drawer-content {
  position: absolute;
  top: 0;
  bottom: 0;
  width: var(--drawer-width);
  background: var(--drawer-background);
  transition: transform var(--duration-normal) var(--ease-spring);
  display: flex;
  flex-direction: column;
}

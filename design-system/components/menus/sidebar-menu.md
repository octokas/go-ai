# Sidebar Menu

```html
<aside class="sidebar-menu">
  <div class="sidebar-header">
    <div class="sidebar-brand">
      <img src="logo.svg" alt="Logo" class="sidebar-logo">
    </div>
    <button class="sidebar-collapse">
      <svg><!-- Collapse icon SVG --></svg>
    </button>
  </div>

  <nav class="sidebar-navigation">
    <div class="sidebar-section">
      <h3 class="sidebar-title">Main Menu</h3>
      <ul class="sidebar-list">
        <li class="sidebar-item">
          <a href="#" class="sidebar-link">
            <svg class="sidebar-icon"><!-- Icon SVG --></svg>
            <span class="sidebar-label">Dashboard</span>
            <span class="sidebar-badge">New</span>
          </a>
        </li>
        <!-- More items -->
      </ul>
    </div>
  </nav>

  <div class="sidebar-footer">
    <div class="sidebar-user">
      <!-- User profile section -->
    </div>
  </div>
</aside>
```

```scss
.sidebar-menu {
  --sidebar-width: 280px;
  --sidebar-collapsed-width: 72px;

  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  width: var(--sidebar-width);
  background: var(--color-surface);
  border-right: 1px solid var(--color-border);
  display: flex;
  flex-direction: column;
  transition: width var(--duration-normal) var(--ease-spring);

  &.collapsed {
    width: var(--sidebar-collapsed-width);

    .sidebar-label,
    .sidebar-badge {
      opacity: 0;
      visibility: hidden;
    }
  }
}

.sidebar-item {
  position: relative;

  &::before {
    content: '';
    position: absolute;
    left: 0;
    width: 3px;
    height: 0;
    background: var(--color-primary);
    transition: height var(--duration-normal) var(--ease-spring);
  }

  &.active::before {
    height: 100%;
  }
}

# Animation Library

## Core Motion Variables
```scss
:root {
  // Easing Functions - Matching our motion system
  --ease-smooth: cubic-bezier(0.2, 0, 0.38, 0.9);
  --ease-smooth-in: cubic-bezier(0.4, 0, 0.2, 1);
  --ease-smooth-out: cubic-bezier(0, 0, 0.2, 1);
  --ease-smooth-in-out: cubic-bezier(0.4, 0, 0.2, 1);
  --ease-smooth-out-in: cubic-bezier(0, 0, 0.2, 1);
  --ease-smooth-bounce: cubic-bezier(0.68, -0.55, 0.265, 1.55);
  --ease-instant: cubic-bezier(0.2, 0, 0, 1);

  // Duration Scale - Matching our motion system
  --duration-instant: 100ms;
  --duration-fast: 150ms;
  --duration-moderate: 200ms;
  --duration-slow: 300ms;
  --duration-slower: 500ms;
}
```

## Transition Mixins
```scss
@mixin smooth-enter {
  opacity: 0;
  transform: scale(1.02);
  transition:
    opacity var(--duration-moderate) var(--ease-smooth),
    transform var(--duration-moderate) var(--ease-smooth);

  &.entered {
    opacity: 1;
    transform: scale(1);
  }
}

@mixin smooth-exit {
  opacity: 1;
  transform: scale(1);
  transition:
    opacity var(--duration-moderate) var(--ease-smooth),
    transform var(--duration-moderate) var(--ease-smooth);

  &.exiting {
    opacity: 0;
    transform: scale(0.98);
  }
}

@mixin smooth-state-change {
  transition: opacity var(--duration-moderate) var(--ease-smooth-in-out);
}
```

## Micro-Interaction Mixins
```scss
@mixin button-interaction {
  transition: transform var(--duration-instant) var(--ease-smooth);

  &:hover {
    transform: scale(1.02);
  }

  &:active {
    transform: scale(0.98);
  }

  &:focus-visible {
    transform: scale(1.02);
  }

  &:disabled {
    opacity: 0.5;
    transform: none;
  }
}

@mixin toggle-interaction {
  transition: transform var(--duration-fast) var(--ease-smooth-bounce);

  &.checked {
    transform: translateX(100%);
  }
}
```

## Gesture Animations
```scss
@mixin gesture-scroll {
  --scroll-distance: 20px;

  @keyframes smoothScroll {
    from {
      transform: translateY(var(--scroll-distance));
      opacity: 0;
    }
    to {
      transform: translateY(0);
      opacity: 1;
    }
  }

  animation: smoothScroll var(--duration-moderate) var(--ease-smooth);
}

@mixin gesture-swipe {
  --swipe-distance: 100%;

  @keyframes smoothSwipe {
    from {
      transform: translateX(var(--swipe-distance));
    }
    to {
      transform: translateX(0);
    }
  }

  animation: smoothSwipe var(--duration-moderate) var(--ease-smooth);
}
```

## Component-Specific Animations
```scss
// Modal
.modal {
  &-enter {
    opacity: 0;
    transform: scale(1.02);
    transition:
      opacity var(--duration-moderate) var(--ease-smooth),
      transform var(--duration-moderate) var(--ease-smooth);
  }

  &-entered {
    opacity: 1;
    transform: scale(1);
  }

  &-exit {
    opacity: 1;
    transform: scale(1);
  }

  &-exiting {
    opacity: 0;
    transform: scale(0.98);
  }
}

// Notification
.notification {
  --notification-offset: 20px;

  &-enter {
    opacity: 0;
    transform: translateY(var(--notification-offset));
    transition:
      opacity var(--duration-fast) var(--ease-smooth),
      transform var(--duration-fast) var(--ease-smooth-bounce);
  }

  &-entered {
    opacity: 1;
    transform: translateY(0);
  }
}

// Drawer
.drawer {
  &-enter {
    transform: translateX(-100%);
    transition: transform var(--duration-slow) var(--ease-smooth);
  }

  &-entered {
    transform: translateX(0);
  }
}
```

## Utility Classes
```scss
.animate {
  &-smooth-enter {
    @include smooth-enter;
  }

  &-smooth-exit {
    @include smooth-exit;
  }

  &-button {
    @include button-interaction;
  }

  &-toggle {
    @include toggle-interaction;
  }

  &-scroll {
    @include gesture-scroll;
  }
}
```

This library now perfectly aligns with our motion system by:
1. Using the same easing functions and timing variables
2. Maintaining the 98%/102% scale ratios for micro-interactions
3. Following the same naming conventions
4. Implementing the same gesture patterns
5. Using consistent transition patterns for enter/exit states

Would you like me to:
1. Add more component-specific animations?
2. Create additional gesture patterns?
3. Add more utility classes?
4. Create animation presets for specific interaction patterns?

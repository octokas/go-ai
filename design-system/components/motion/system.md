# Motion System

## Timing Functions
```scss
:root {
  // Natural Movement
  --ease-default: cubic-bezier(0.2, 0, 0.38, 0.9);
  --ease-spring: cubic-bezier(0.175, 0.885, 0.32, 1.275);
  --ease-bounce: cubic-bezier(0.68, -0.55, 0.265, 1.55);

  // Entrance & Exit
  --ease-enter: cubic-bezier(0, 0, 0.38, 0.9);
  --ease-exit: cubic-bezier(0.2, 0, 1, 0.9);

  // Duration Scale
  --duration-instant: 100ms;
  --duration-fast: 200ms;
  --duration-normal: 300ms;
  --duration-slow: 400ms;
  --duration-slower: 500ms;
}

// Motion Mixins
@mixin motion-fade {
  &.enter { animation: fade-in var(--duration-normal) var(--ease-enter); }
  &.exit { animation: fade-out var(--duration-normal) var(--ease-exit); }
}

@mixin motion-slide {
  &.enter { animation: slide-in var(--duration-normal) var(--ease-spring); }
  &.exit { animation: slide-out var(--duration-normal) var(--ease-exit); }
}

@mixin motion-scale {
  &.enter { animation: scale-in var(--duration-normal) var(--ease-spring); }
  &.exit { animation: scale-out var(--duration-normal) var(--ease-exit); }
}

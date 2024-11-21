# Motion System

Our motion system creates fluid, purposeful animations that enhance usability while maintaining that native feel of macOS.

## Timing & Easing Functions
```scss
// Easing
--smooth: cubic-bezier(0.2, 0, 0.38, 0.9); // Smooth and predictable
--smooth-in: cubic-bezier(0.4, 0, 0.2, 1); // Starts slow, ends fast
--smooth-out: cubic-bezier(0, 0, 0.2, 1); // Starts fast, ends slow
--smooth-in-out: cubic-bezier(0.4, 0, 0.2, 1); // Starts slow, ends fast, then starts fast, ends abruptly
--smooth-out-in: cubic-bezier(0, 0, 0.2, 1); // Starts slow, ends fast, then starts fast, ends abruptly
--smooth-in-out-in: cubic-bezier(0.4, 0, 0.2, 1); // Starts fast, ends abruptly, then starts slow, ends fast
--smooth-out-in-out: cubic-bezier(0, 0, 0.2, 1); // Starts slow, ends fast, then starts fast, ends abruptly

// Special Effects
--instant: cubic-bezier(0.2, 0, 0, 1); // Instant, no easing
--smooth-bounce: cubic-bezier(0.68, -0.55, 0.265, 1.55); // Smooth bounce, creates a bouncing effect
```

### Duration Tokens
```scss
// milliseconds
--instant: 100ms; // Similar to a click
--fast: 150ms; // Similar to a hover
--moderate: 200ms; // Similar to a modal transition
--slow: 300ms; // Similar to a page transition
--slower: 500ms; // Similar to a modal transition
```

### Motion Patterns

#### Transitions
```scss
// 98% and 102% are the magic numbers for the scale down and scale up
// https://www.youtube.com/watch?v=jXU9HGuMhGI

// Transitions
--smooth-enter: fade-in + scale(1.02); // Fade in + Scale up slightly, 102%
--smooth-exit: fade-out + scale(0.98); // Fade out + Scale down slightly, 98%
--smooth-active: pulse(scale(1.02) and scale(0.98)); // Bounce slightly (102% and 98%)
--smooth-hover: fade-in + scale(1.02); // Fade in + Scale up slightly, 102%
--smooth-focus: fade-in + scale(1.02); // Fade in + Scale up slightly, 102%
--smooth-state-change: smooth crossfade; // Smooth crossfade
--smooth-paused: fade-out; // Fade out
```

#### Micro Interactions
```scss
// Button are same for Toggle, Slider, Textbox, Select, Checkbox
// 98% and 102% are the magic numbers for the scale down and scale up
// https://www.youtube.com/watch?v=jXU9HGuMhGI

// Button
--button-press: scale(0.98); // Scale down slightly, 98%
--button-release: scale(1.02); // Scale up slightly, 102%
--button-hover: fade-in + scale(1.02); // Fade in + Scale up slightly, 102%
--button-focus: fade-in + scale(1.02); // Fade in + Scale up slightly, 102%
--button-active: pulse(scale(1.02) and scale(0.98)); // Bounce slightly (102% and 98%)
--button-inactive: fade-out; // Fade out

// Checkbox, Slider, Toggle
--checkbox-toggle: bounce(0.02); // Bounce slightly
--slider-toggle: bounce(0.02); // Bounce slightly
--toggle-disabled: fade-out; // Fade out
```

#### Gestures
```scss
// Sweeping Gestures
--gesture-scroll: smooth crossfade; // Smooth crossfade, natural scroll
--gesture-swipe: smooth crossfade; // Smooth crossfade, natural swipe
--gesture-two-finger-scroll: smooth crossfade + scale(1.02); // Smooth crossfade + Scale up slightly, 102%, momentum scroll

// Tapping Gestures
--gesture-tap: fade-in + scale(1.02); // Fade in + Scale up slightly, 102%, like a tap
--gesture-long-press: fade-in + scale(1.02); // Fade in + Scale up slightly, 102%, holding a tap
--gesture-press: scale(0.98); // Scale down slightly, 98%, pressing a button
--gesture-release: scale(1.02); // Scale up slightly, 102%, releasing a button

// Hovering Gestures
--gesture-hover: fade-in + scale(1.02); // Fade in + Scale up slightly, 102%, hovering over a button
--gesture-focus: fade-in + scale(1.02); // Fade in + Scale up slightly, 102%, focusing on a button
--gesture-active: pulse(scale(1.02) and scale(0.98)); // Bounce slightly (102% and 98%), activating a button

// Disabled Gestures
--gesture-disabled: fade-out; // Fade out
```

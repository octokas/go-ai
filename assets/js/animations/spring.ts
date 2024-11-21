interface SpringConfig {
  stiffness: number;
  damping: number;
  mass: number;
}

// Apple-like spring configurations
const SPRING_CONFIGS = {
  // Regular UI interactions
  default: {
    stiffness: 300,
    damping: 30,
    mass: 1,
  },
  // Quick, snappy responses
  responsive: {
    stiffness: 400,
    damping: 40,
    mass: 1,
  },
  // Smooth, bouncy animations
  bouncy: {
    stiffness: 200,
    damping: 20,
    mass: 1,
  },
} as const;

class SpringAnimation {
  private spring: SpringConfig;
  private target: number;
  private current: number;
  private velocity: number;

  constructor(config: SpringConfig) {
    this.spring = config;
    this.target = 0;
    this.current = 0;
    this.velocity = 0;
  }

  animate(callback: (value: number) => void) {
    const animate = () => {
      const force = -this.spring.stiffness * (this.current - this.target);
      const damping = -this.spring.damping * this.velocity;
      const acceleration = (force + damping) / this.spring.mass;

      this.velocity += acceleration * (1 / 60); // 60fps
      this.current += this.velocity * (1 / 60);

      callback(this.current);

      if (Math.abs(this.target - this.current) > 0.001) {
        requestAnimationFrame(animate);
      }
    };

    requestAnimationFrame(animate);
  }
}

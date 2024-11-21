import { Controller } from "@hotwired/stimulus"
import { SpringAnimation } from "../animations/spring"

export default class extends Controller {
    static targets = ["item"]

    connect() {
        this.setupAnimations()
    }

    async createTask(event) {
        const response = await fetch("/api/tasks", {
            method: "POST",
            body: JSON.stringify(event.detail)
        })

        if (response.ok) {
            const task = await response.json()
            this.animateNewTask(task)
        }
    }

    private animateNewTask(task) {
        const template = document.getElementById("task-item-template")
        const clone = template.content.cloneNode(true)

        // Populate template with task data

        const spring = new SpringAnimation(SPRING_CONFIGS.bouncy)
        clone.style.opacity = "0"
        clone.style.transform = "translateY(20px)"

        this.element.appendChild(clone)

        spring.animate((value) => {
            clone.style.opacity = String(value)
            clone.style.transform = `translateY(${20 * (1 - value)}px)`
        })
    }
}

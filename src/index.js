require('./mystyles.scss');

import Alpine from 'alpinejs'

window.Alpine = Alpine

Alpine.data('app', () => ({
    minWidth: 768, // Minimum width for desktop devices
    fadeElementList: ["sobreMim", "abordagem", "academico", "localizacao"],
    init() {
        this.showElement(window.scrollY)
    },    
    get isMobile() {
        return 'ontouchstart' in window || navigator.maxTouchPoints > 0
    },    
    showElement(scrollY) {
        this.fadeElementList.forEach(element => {
            const el = this.$refs[element]
            if (this.isMobile || el.getBoundingClientRect().top <= scrollY && el.getBoundingClientRect().bottom >= scrollY)
                el.style.opacity = 1            
        })
    },
    mudaCorLink(event, cor) { 
        event.target.closest('a').classList.add(cor)
    },
    removeCorLink(event, cor) { 
        event.target.closest('a').classList.remove(cor)
    }
}))

Alpine.start()
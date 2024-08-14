require('./mystyles.scss');

import Alpine from 'alpinejs'

window.Alpine = Alpine

Alpine.data('app', () => ({
    minWidth: 768, // Minimum width for desktop devices
    fadeElementList: ["logo", "sobreMim", "abordagem", "academico", "localizacao"],
    imagesLoaded: {},
    init() {
        this.loadImages().then(() => {
            this.showElement();
        });
    },
    async loadImages() {
        for (let ref of this.fadeElementList) {
            const el = this.$refs[ref]
            const images = el.querySelectorAll('img')
            for (let img of images) {
                await this.loadImage(img)
                this.imagesLoaded[img.src] = true
            }
        }
    },
    loadImage(img) {
        return new Promise((resolve) => {
            img.onload = resolve
            img.onerror = () => {
                console.error('Erro ao carregar a imagem:', img.src);
                resolve(); // Mesmo que a imagem não carregue, resolva a promise
            }
            img.src = img.dataset.src
        })
    },
    get isMobile() {
        return 'ontouchstart' in window || navigator.maxTouchPoints > 0
    },
    showElement() {
        this.fadeElementList.forEach(ref => {
            const el = this.$refs[ref]
            const images = el.querySelectorAll('img')
            const isInViewport = el.getBoundingClientRect().top < window.innerHeight && el.getBoundingClientRect().bottom >= 0
            if (this.isMobile || isInViewport) {
                if (Array.from(images).every(img => this.imagesLoaded[img.src])) {
                    el.classList.add('show')
                }
            }
        })
    }
}))

Alpine.start()

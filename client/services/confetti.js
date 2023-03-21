import confetti from 'canvas-confetti';
import moment from 'moment-timezone';

const launchConfetti = (isShiny, isFirstTry) => {
    var duration = 3 * 1000;
    var animationEnd = moment() + duration;
    var particleCount = 200;
    var defaults = {
        startVelocity: 25,
        spread: 360,
        ticks: 300,
        zIndex: 0,
        scalar: 1.4,
    };

    if (isShiny || isFirstTry) {
        defaults.shapes = ['star'];
        defaults.scalar = 1.1;
        particleCount = 90;
        if (isShiny && isFirstTry) {
            defaults.colors = ['9EFD38', '32CD32', 'A8E4A0', '98FB98', '7CFC00'];
        } else if (isShiny) {
            defaults.colors = ['FFE400', 'FFBD00', 'E89400', 'FFCA6C', 'FDFFB8'];
        } else if (isFirstTry) {
            defaults.colors = ['00E5EE', '00FFFF', 'E0FFFF', '98F5FF'];
        }
    }

    const randomInRange = (min, max) => {
        return Math.random() * (max - min) + min;
    };

    const fire = () => {
        var timeLeft = animationEnd - moment();

        if (timeLeft <= 0) {
            return clearInterval(launchConfetti);
        }

        var currParticleCount = particleCount * (timeLeft / duration);
        // since particles fall down, start a bit higher than random
        confetti(
            Object.assign({}, defaults, {
                currParticleCount,
                origin: {
                    x: randomInRange(0.1, 0.3),
                    y: Math.random() - 0.4,
                },
            })
        );
        confetti(
            Object.assign({}, defaults, {
                currParticleCount,
                origin: {
                    x: randomInRange(0.7, 0.9),
                    y: Math.random() - 0.4,
                },
            })
        );
    };

    fire();

    setInterval(() => {
        fire();
    }, 250);
};

export { launchConfetti };

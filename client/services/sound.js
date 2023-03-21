
const playWinnerSound = () => {
    const audioEmeraldWinMusic = new Audio('/client/assets/audio/emerald-winner.mp3');
    const audioFireworks = new Audio('/client/assets/audio/fireworks.mp3');

    audioEmeraldWinMusic.volume = 0.2;
    audioFireworks.volume = 0.15;

    audioEmeraldWinMusic.play();
    audioFireworks.play();
};

export {
    playWinnerSound,
}
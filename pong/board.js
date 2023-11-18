initializeLines();

function initializeLines() {
    console.log('drawing lines');
    const game = document.getElementById('game');
    const ctx = game.getContext('2d');
    ctx.beginPath();
    ctx.moveTo(400, 0);
    ctx.lineTo(400,600);
    ctx.lineWidth = 1;
    ctx.strokeStyle = "white";
    ctx.stroke();
    console.log('Lines drawn!');

    //draw paddles
    drawLeftPaddle(ctx);
    drawRightPaddle(ctx);

    //draw ball
    drawBall(ctx);


}

function drawBall(context){
    context.beginPath();
    context.arc(300, 100, 10, 0, 2 * Math.PI);
    context.strokeStyle = "green";
    context.stroke();
}

function drawLeftPaddle(context) {
    context.beginPath();
    context.rect(20, 300, 5, 80);
    context.strokeStyle = "#ADD8E6";
    context.stroke();
}

function drawRightPaddle(context) {
    context.beginPath();
    context.rect(780, 300, 5, 80);
    context.strokeStyle = "#ADD8E6";
    context.stroke();
}


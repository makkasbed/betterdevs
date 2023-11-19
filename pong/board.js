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
    drawPaddle(ctx, 20, 300, 5, 80, "human"); //left paddle
    drawPaddle(ctx, 780, 300, 5, 80, "computer"); // right paddle


    //draw ball
    drawBall(ctx);


}

function drawPaddle(context,width,height,xpos, ypos, type){
    context.beginPath();
    context.rect(width, height, xpos, ypos);
    context.strokeStyle = "#ADD8e6";
    context.stroke();
}


function drawBall(context){
    context.beginPath();
    context.arc(300, 100, 10, 0, 2 * Math.PI);
    context.strokeStyle = "green";
    context.stroke();
}
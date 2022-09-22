const {Input, Output} = require('./calculator/calculator_pb.js');
const {CalculatorServiceClient} = require('./calculator/calculator_grpc_web_pb.js');

// here 8080 is envoy port
// envoy will forward to grpc server and respond to client
var client = new CalculatorServiceClient('http://localhost:8080');

const squareInput = document.getElementById('square-input');
const squareFind = document.getElementById('square-button')
squareFind.addEventListener('click', () => {
    var input = new Input();
    input.setNumber(squareInput.value);
    client.findSquare(input, {}, (err, r) => {
        addResponse(r);
    });
});

const factorInput = document.getElementById('factor-input');
const factorFind = document.getElementById('factor-button');

factorFind.addEventListener('click', () => {
    var input = new Input();
    input.setNumber(factorInput.value);
    var stream = client.findFactors(input, {});
    stream.on('data', (r) => {
        addResponse(r);
    });
    stream.on('status', (status) => {
        console.log(status.code);
    });
    stream.on('end', (end) => {

    });
    document.getElementById('factor-stop').addEventListener('click', () => {
        stream.cancel();
    });
});
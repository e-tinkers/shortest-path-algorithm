# Shortest Path Algorithm

This repo contains 3 different implementation of shortest path algorithm in Python, Nodejs(JavaScript) and Golang(Go). The code included some test code for benchmark purpose, it also serves as an example on how to use the algorithm. For detail disussion on actual application of the algorithm, you could take a look at my blog post [How to create an interactive transport system map with shortest path algorithm](https://www.e-tinkers.com/2018/07/how-to-create-an-interactive-transport-system-map-with-shortest-path-algorithm/). The start and end nodes in the test code are two actual stations of Singapore Mass Rapid Transit (SMRT) system which can be visually seen here:

![The shortest path between ns1/ew24 and ew2/dt32](./images/logo.png)

To run the Golang test code:

    go run shortest.go

To run the Nodejs test code:

    node shortest.js

To run the Python test code:

    python3 shortest.py

## Performanace

The code was originally written in Python with the intention for running it on a Raspberry Pi 3B, the code was too slow to be used practically, and port to JavaScript and running on the client side with the assumption that any computer or even mobile phone nowadays is faster than a Raspberry Pi. I recently port it into Golang and run some benchmark on both Raspberry Pi and my MacBook Pro late 2011 model, here are the comparison of the test results.

<table style="table-layout: fixed; width: 740px">
<colgroup>
<col style="width: 140px">
<col style="width: 100px">
<col style="width: 100px">
<col style="width: 100px">
<col style="width: 100px">
<col style="width: 100px">
<col style="width: 100px">
</colgroup>
  <tr>
    <th rowspan="2"></th>
    <th colspan="3">Raspberry Pi 3 model B</th>
    <th colspan="3">MacBook Pro</th>
  </tr>
  <tr>
    <td>go1.12.6</td>
    <td>node v8.11.1</td>
    <td>python3.5.3</td>
    <td>go1.12.1</td>
    <td>node v10.13.0</td>
    <td>python3.6.1</td>
  </tr>
  <tr>
    <td rowspan="10">Time (sec)</td>
    <td>1.968</td>
    <td>6.568</td>
    <td>16.653</td>
    <td>0.382</td>
    <td>1.209</td>
    <td>2.314</td>
  </tr>
  <tr>
    <td>2.072</td>
    <td>6.624</td>
    <td>16.711</td>
    <td>0.392</td>
    <td>1.210</td>
    <td>2.159<br></td>
  </tr>
  <tr>
    <td>1.969</td>
    <td>6.576</td>
    <td>16.619</td>
    <td>0.394</td>
    <td>1.189</td>
    <td>2.363</td>
  </tr>
  <tr>
    <td>2.021</td>
    <td>6.582</td>
    <td>16.650</td>
    <td>0.383</td>
    <td>1.164</td>
    <td>2.255</td>
  </tr>
  <tr>
    <td>1.965</td>
    <td>6.602</td>
    <td>16.751</td>
    <td>0.388</td>
    <td>1.188</td>
    <td>2.262</td>
  </tr>
  <tr>
    <td>1.990</td>
    <td>6.614</td>
    <td>16.656</td>
    <td>0.380</td>
    <td>1.217</td>
    <td>2.285</td>
  </tr>
  <tr>
    <td>2.031</td>
    <td>6.592</td>
    <td>16.573</td>
    <td>0.407</td>
    <td>1.186</td>
    <td>2.235</td>
  </tr>
  <tr>
    <td>1.964</td>
    <td>6.562</td>
    <td>16.854</td>
    <td>0.401</td>
    <td>1.195</td>
    <td>2.221</td>
  </tr>
  <tr>
    <td>1.957</td>
    <td>6.645</td>
    <td>16.889</td>
    <td>0.408</td>
    <td>1.187</td>
    <td>2.205</td>
  </tr>
  <tr>
    <td>1.962</td>
    <td>6.562</td>
    <td>16.624</td>
    <td>0.397</td>
    <td>1.161</td>
    <td>2.201</td>
  </tr>
  <tr>
    <td>Average (sec)</td>
    <td>1.990</td>
    <td>6.593</td>
    <td>16.698</td>
    <td>0.393</td>
    <td>1.191</td>
    <td>2.250</td>
  </tr>
</table>

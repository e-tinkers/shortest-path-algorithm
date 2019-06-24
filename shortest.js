var fs = require('fs');

function calShortestPath(graph, start, end, path=[]) {
  if (!graph.hasOwnProperty(start)) {
    return [];
  }
  path = path.concat(start);
  if (start == end) {
    return path;
  }
  let shortest = [];
  for (let node of graph[start]) {
    if (path.indexOf(node) == -1) {
      let newPath = calShortestPath(graph, node, end, path);
      if (newPath.length > 0) {
        if (shortest.length == 0 || newPath.length < shortest.length) {
          shortest = newPath;
        }
      }
    }
  };
  return shortest;
}

var s = "ns1/ew24";
var e = "ew2/dt32";
var p = [];
var graph = JSON.parse(fs.readFileSync('./data/stations_sg.json', 'utf8'));

const t = new Date();
var sp = calShortestPath(graph, s, e, p);
const elapsed = new Date() - t;
console.log(elapsed);
console.log(sp);

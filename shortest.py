import json
import time

def load_graph(file_name):
    with open(file_name) as f:
        data = json.load(f)
        return data


def shortest_path(graph, start, end, path=[]):
    path = path + [start]
    if start == end:
        return path
    if start not in graph.keys():
        return None
    shortest = None
    for node in graph[start]:
        if node not in path:
            new_path = shortest_path(graph, node, end, path)
            if new_path:
                if not shortest or len(new_path) < len(shortest):
                    shortest = new_path
    return shortest


stations = load_graph('./data/stations_sg.json')
s = time.time()
route = shortest_path(stations, "ns1/ew24", "ew2/dt32")
elapsed = time.time() - s
print(elapsed)
print(route)

# https://adventofcode.com/2023/day/8
# would need to donwload the file from the link above

import re
import math

def processGraph(graph: list[str]) -> dict[str, tuple[str, str]]:
    mapping = {}
    f_string = "([0-9A-Z]{3}) = \(([0-9A-Z]{3}), ([0-9A-Z]{3})\)"
    for l in graph[:-1]:
        parsed = re.match(f_string, l)
        if parsed:
            node, left, right = parsed.group(1, 2, 3)
            mapping[node] = (left, right)
        else:
            print(l)
            raise Exception
    return mapping

def solve1(directions, map):
    node = "AAA"
    steps = 0
    while True:
        for d in directions:
            next = 0 if d == "L" else 1
            node = map[node][next]
            
            steps += 1
            if node == "ZZZ":
                print("solution 1", steps)
                return

def solve2(directions, map):
    nodes = [m for m in map if m[2] == 'A']
    node_steps = []
    
    for n in nodes:
        steps = 0
        current = n
        instructionIndex = 0
        while current[-1] != "Z":
            if instructionIndex == len(directions):
                instructionIndex = 0
            instruction = 1 if "R" == directions[instructionIndex] else 0
            current = map[current][instruction]
            instructionIndex += 1
            steps += 1 
        node_steps.append(steps)
    print("solution 2", math.lcm(*node_steps))

with open('input8.txt', 'r') as f:
    input = f.read()
    lines = input.split('\n')

    directions = lines[0]
    graph = lines[2:]
    map = processGraph(graph)

solve1(directions, map)
solve2(directions, map)





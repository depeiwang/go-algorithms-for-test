import fs from "fs";

type Point = [number, number] // x, y
type Curve = Array<Point>

function curvePoints(curve: Curve): Set<string> {
    const s = new Set<string>;
    for (const p of curve) {
        const pStr = p[0] + "," + p[1]
        s.add(pStr);
    }
    return s;
}

function hasIntersection(curve1Points: Set<string>, curve2Points: Set<string>): boolean {
    for (const ps of curve1Points.keys()) {
        if (curve2Points.has(ps)) {
            return true;
        }
    }
    return false;
}

function findMin(arr: Array<Curve>, index: 0 | 1): number {
    let min = Infinity
    for (const curve of arr) {
        for (const point of curve) {
            if (point[index] < min) {
                min = point[index]
            }
        }
    }

    return min;
}

function findMax(arr: Array<Curve>, index: 0 | 1): number {
    let max = -Infinity
    for (const curve of arr) {
        for (const point of curve) {
            if (point[index] > max) {
                max = point[index]
            }
        }
    }
    return max;
}

// 判断所有曲线是否相连
function allCurvesConnected(arr: Array<Curve>): boolean {
    const setList: Array<Set<string>> = [];
    for (const curve of arr) {
        setList.push(curvePoints(curve));
    }

    for (let i = 0; i < setList.length - 1; i++) {
        let connected = false;
        for (let j = i + 1; j < setList.length; j++) {
            connected = hasIntersection(setList[i], setList[j]);
            if (connected) {
                break
            }
        }
        if (!connected) {
            return false;
        }
    }
    return true;
}

// 寻找边界
function findEdge(arr: Array<Curve>): { topLeft: Point, bottomRight: Point } {
    return {
        topLeft: [findMin(arr, 0), findMin(arr, 1)],
        bottomRight: [findMax(arr, 0), findMax(arr, 1)]
    }
}

function main() {
    const str = fs.readFileSync("./good3.json", { encoding: 'utf-8' });
    const data = JSON.parse(str) as { data: Array<Curve> };
    const connected = allCurvesConnected(data.data);
    console.log("connected: ", connected);

    const edge = findEdge(data.data);
    console.log("edge:", edge);

}

main();

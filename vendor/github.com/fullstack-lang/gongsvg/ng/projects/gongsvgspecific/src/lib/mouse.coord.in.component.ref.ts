import * as gongsvg from 'gongsvg'
import { createPoint } from './link/draw.segments'

export function mouseCoordInComponentRef(event: MouseEvent): gongsvg.PointDB {

    const targetElement = event.target as SVGElement
    const svg = targetElement.ownerSVGElement

    if (svg) {
        const point = svg.createSVGPoint()

        point.x = event.clientX;
        point.y = event.clientY;

        // console.log(point.x, point.y); // window x, y

        const SVGmatrix = svg.getScreenCTM()

        const localPoint = point.matrixTransform(SVGmatrix!.inverse());
        let res = createPoint(localPoint.x, localPoint.y)
        return res
    } else {
        console.log("mouseCoorInComponentRef: no svg found in event")
        let res = createPoint(0, 0)
        return res
    }
}
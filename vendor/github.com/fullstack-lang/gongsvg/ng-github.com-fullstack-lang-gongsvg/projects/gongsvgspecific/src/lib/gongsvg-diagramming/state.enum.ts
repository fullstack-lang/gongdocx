// StateEnumType are states supported by the drawing-svg component
export enum StateEnumType {
    NOT_EDITABLE = "NOT_EDITABLE",

    WAITING_FOR_USER_INPUT = "WAITING_FOR_USER_INPUT",

    MULTI_RECTS_SELECTION = "MULTI_RECTS_SELECTION",
    LINK_DRAWING = "LINK_DRAWING",

    RECTS_DRAGGING = "RECTS_DRAGGING",
    RECT_ANCHOR_DRAGGING = "RECT_ANCHOR_DRAGGING",
    LINK_ANCHORED_TEXT_DRAGGING = "LINK_ANCHORED_TEXT_DRAGGING",
    LINK_DRAGGING = "LINK_DRAGGING",
}
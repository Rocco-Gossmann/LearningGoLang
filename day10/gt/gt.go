package gt

// TickHandlers take a float32 as an argument and return a boolean
type TickHandler func(timeMultiplier float32) bool

const handlerListSizeLimit = 32
var handlerList1 = make([]TickHandler, 0, handlerListSizeLimit);
var handlerList2 = make([]TickHandler, 0, handlerListSizeLimit);

var currentHandlerList *[]TickHandler = &handlerList1;
var nextHandlerList    *[]TickHandler = &handlerList2;

const capacyityReachedMsg = "capacity reached";

func RegisterHandler(handler TickHandler) {
    if handler != nil {
        if len(*nextHandlerList) < handlerListSizeLimit {
            *nextHandlerList = append(*nextHandlerList, handler);

        } else {
            panic(capacyityReachedMsg);
        }
    }
}

func HandleTick() int {

    const timeMultiplyer = 0; // <-- this does not matter right now

    // Reslice the current list to have no elements
    *currentHandlerList = (*currentHandlerList)[0:0] 

    // Swap the contents of the two lists: CurrentList == filled, nextList == empty
    currentHandlerList, nextHandlerList = nextHandlerList, currentHandlerList

    for _, handler := range *currentHandlerList {
        // a handler responds with true, mark it for the next round
        if handler(timeMultiplyer) {
            *nextHandlerList = append(*nextHandlerList, handler);
        }
    }

    return len(*nextHandlerList);
}

package state

import (
    "log"
    "time"
)

type State struct {
    Url   string
    State string
}

func logState(stateMap map[string]State) {
    log.Println("\"Monitor\": Current state:")
    for _, state := range stateMap {
        log.Printf("    %s:     %s", state.Url, state.State)
    }
    log.Println("-------------------------")
}

// Monitor 会返回一个基于State的只写通道，你可以获取该通道，并且每次有新State更新时，向通道中更新数据
// 该方法会根据lookInterval时间来打印当前所有的State状态
// 注意，该方法只需要执行一次即可，并不需要每次执行来获取一个新的通道
// 返回的stateMap可以用于其他用途
func Monitor(lookInterval time.Duration) (chan<- State, map[string]State) {
    updateChan := make(chan State)
    stateMap := make(map[string]State)
    log.Printf("Monitor starts listening the change of stateChan and print state of stateMap in circle.\n")
    go func() {
        ticker := time.NewTicker(lookInterval)
        for {
            select {
            case <-ticker.C:
                log.Printf("timer----\n")
                logState(stateMap)
            case updateState := <-updateChan:
                // update stateMap
                if value, ok := stateMap[updateState.Url]; !ok || value != updateState {
                    log.Printf("Monitor update the content of %q, value: %q.\n", updateState.Url, updateState.State)
                    stateMap[updateState.Url] = updateState
                    // If updating, try to print state.
                }
            }
        }
    }()

    return updateChan, stateMap
}

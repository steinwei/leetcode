export class logger<T>{
    info(...arg:T[]){
        console.log(...arg)
    }

    time(...arg:T[]){
        const now = Date.now()
        console.log(now)
        return now
    }
}

// export 
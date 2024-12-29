import {jobUpdateStreamViewCount} from "./sendStreamViewCount";


export const  setupCronJobs =() => {
    jobUpdateStreamViewCount.start()
}
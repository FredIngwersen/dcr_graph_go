package main
import(
    "fmt"
)
func test_graph_1(traces map[string][]string)[2]int{
    result := [2]int{0,0}
    i := 0
    for _, trace := range traces  {
        if(trace[0] != "Fill out application"){
            fmt.Println(trace[0])
            fmt.Println(i)
            result[1] += 1
        }else{
            result[0] += 1
        }
        i++
    }
    return result
}

func test_graph_2(traces map[string][]string)[2]int{

    result := [2]int{0,0}
    number_of_false := 0

    for _, trace := range traces {
        applicant_informed := true
        change_phase := true

        for _, event := range trace{
            if(event == "Reject"){
                applicant_informed = false
                change_phase = false
            }else if(event == "Applicant informed"){
                applicant_informed = true
            }else if(event == "Change phase to Abort"){
                change_phase = true
            }
        }
        if(!(applicant_informed && change_phase)){
            number_of_false += 1
        }
        result[0] = len(traces) - number_of_false
        result[1] = number_of_false
    }
    return result
}

func test_graph_3(traces map[string][]string)[2]int{
    result := [2]int{0,0}
    number_of_false := 0

    for _, trace := range traces {
        first_payment := false
        for _, event := range trace{
            if(event == "First Payment" && first_payment){
                number_of_false += 1
                break
            }else if(event == "First Payment"){
                first_payment = true
            }
        }
        result[0] = len(traces) - number_of_false
        result[1] = number_of_false
    }
    return result
}
func test_graph_4(traces map[string][]string)[2]int{
    result := [2]int{0,0}
    number_of_false := 0
    for _, trace := range traces{
        lawyer_review := false
        architect_review := false
        for _, event := range trace{
        if((event == "Architect Review" && lawyer_review) ||
            (event == "Lawyer Review" && architect_review)){
                number_of_false += 1
                break
        }else if(event == "Architect Review"){
            architect_review = true
        }else if(event =="Lawyer Review"){
            lawyer_review = true
        }
    }
    result[0] = len(traces) - number_of_false
    result[1] = number_of_false
    }
    return result
}

func main(){
    csv_file := "log.csv"
    traces := get_traces(csv_file, ";")
    fmt.Println(test_graph_1(traces))
    fmt.Println(test_graph_2(traces))
    fmt.Println(test_graph_3(traces))
    fmt.Println(test_graph_4(traces))
}

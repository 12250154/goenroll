window.onload = function () {
    fetch('/student/all')
    .then(response => response.text())
    .then(data => showStudents(data))
}

function newRoW(table,student){
    var row = table.insertRow(table.length)
    var td = []
    for (i=0; i<table.rows[0].cells.length; i++){
           td[i] = row.insertCell(i)
    }
        //insert data in the td cells
    td[0].innerHTML = student.stdid
    td[1].innerHTML = student.fname
    td[2].innerHTML = student.lname
    td[3].innerHTML = student.email
    td[4].innerHTML = '<input type="button" onclick="deleteStudent(this)" value="delete" id="button-1" />'
    td[5].innerHTML = '<input type="button" onclick="updateStudent(this)" value="edit" id="button-2" />'
}

function showStudents(data){
    var students = JSON.parse(data)
    var table = document.getElementById("myTable")
    students.forEach(stud => {
    newRoW(table,stud) 
    })
}

function showStudent(data){
    //console.log(data)
    //convert json string to js obj
    const student = JSON.parse(data)
    var table = document.getElementById("myTable")
    newRoW(table,student)
}

//helper function to reset form fields
function resetform(){
    document.getElementById("sid").value = ""
    document.getElementById("fname").value = ""
    document.getElementById("lname").value = ""
    document.getElementById("email").value = ""
}

//helper function to get form data
function getFormData() {
    //create a javascript object to store form data
    var formData = {
        stdid : parseInt(document.getElementById("sid").value),
        fname : document.getElementById("fname").value,
        lname : document.getElementById("lname").value,
        email : document.getElementById("email").value
    }
    return formData
}

function addStudent(){
    let formdata = getFormData()
    //form validation
    var sid = formdata.stdid
    if (isNaN(sid)){
        alert("Enter valid student ID")
        return
    } else if (formdata.email == ""){
        alert("Email cannot be empty")
        return
    } else if(formdata.fname ==""){
        alert("First name cannot be empty")
        return
    }

    //call POST API
    //axios, fetch to make http request
    //API route, request obj
    fetch('/student/add', {
        method: "POST",
        body: JSON.stringify(formdata),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(response1 => { //storing the response of response1
        //check the response from fetch is resolve or rejected
        if (response1.ok) {
            // /student/1001
            fetch("/student/"+sid) 
            .then(response2 => response2.text())  //then responsible for handling the promise
            .then(data => showStudent(data))
        } else {
            throw new Error(response1.status)
        }
    }).catch(e => {
        if (e.message == 401){
            alert("User Not logged in")
            window.open("index.html", "_self")
        } else if (e.message == 400){
            alert("Bad Request")
        } else {
            alert("Internal Server Error")
        }
    });
    resetform();
} 

function updateAPIRequest(oldSid) {
    //get the updated data from the user
    let newdata = getFormData()
    //call update API
    fetch("/student/"+oldSid,{
        method: "PUT",
        body: JSON.stringify(newdata),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(res =>{
        if (res.ok){
            //fill in selected row with updated value
            selectedRow.cells[0].innerHTML = newdata.stdid
            selectedRow.cells[1].innerHTML = newdata.fname
            selectedRow.cells[2].innerHTML = newdata.lname
            selectedRow.cells[3].innerHTML = newdata.email

            //change the button value to initial state
            var btn = document.getElementById("button-add")
            btn.innerHTML = "Add"
            btn.setAttribute("onclick", "addStudent()")

            selectedRow = null
            resetform()
        } else{
            alert("server: update request error")
        }
    })
}
var selectedRow =null
function updateStudent(input){
    //get the selected row
    selectedRow = input.parentElement.parentElement
    document.getElementById("sid").value = selectedRow.cells[0].innerHTML
    document.getElementById("fname").value = selectedRow.cells[1].innerHTML
    document.getElementById("lname").value = selectedRow.cells[2].innerHTML
    document.getElementById("email").value = selectedRow.cells[3].innerHTML
    
    sid = selectedRow.cells[0].innerHTML

    //change button value to update
    var btn = document.getElementById("button-add")
    btn.innerHTML = "update"
    btn.setAttribute("onclick", "updateAPIRequest(sid)")
}

function deleteStudent(input){
    if (confirm("Are you sure to Delete this?")){
        selectedRow = input.parentElement.parentElement
        sid = selectedRow.cells[0].innerHTML
        fetch("/student/"+sid,{
            method: "DELETE"
        })
        .then(res => {
            if(res.ok){
                var rowIndex = selectedRow.rowIndex
                document.getElementById("myTable").deleteRow(rowIndex)
                selectedRow = null
            } else{
                alert("Server: delete request error")
            }
        })
    }
}


// ================= LOAD COURSES WHEN PAGE OPENS =================
const API = "http://localhost:8080";
window.onload = function () {
    fetch('/course/all')
    .then(res => res.json())
    .then(data => showCourses(data))
}

// ================= CREATE TABLE ROW =================
function newRow(table, course){
    let row = table.insertRow(table.length)
    let td = []

    for (let i = 0; i < table.rows[0].cells.length; i++){
        td[i] = row.insertCell(i)
    }

    td[0].textContent = course.courseid
    td[1].textContent = course.coursename

    td[2].innerHTML = '<input type="button" onclick="deleteCourse(this)" value="delete" id="button-1"/>'
    td[3].innerHTML = '<input type="button" onclick="editCourse(this)" value="edit" id="button-2"  />'
}

// ================= SHOW ALL COURSES =================
function showCourses(courses){
    let table = document.getElementById("myTable")
    courses.forEach(course => newRow(table, course))
}

// ================= SHOW SINGLE COURSE AFTER ADD =================
function showCourse(course){
    let table = document.getElementById("myTable")
    newRow(table, course)
}

// ================= RESET FORM =================
function resetForm(){
    document.getElementById("cid").value = ""
    document.getElementById("cname").value = ""
}

// ================= GET FORM DATA =================
function getCourseFormData(){
    let formData = {
        courseid: parseInt(document.getElementById("cid").value),
        coursename: document.getElementById("cname").value
    }
    return formData
}

// ================= ADD COURSE =================
function addCourse(){

    let data = getCourseFormData()

    if(data.courseid == ""){
        alert("Enter course ID")
        return
    }
    if(data.coursename == ""){
        alert("Enter course name")
        return
    }

    fetch('/course/add', {
        method: "POST",
        credentials: "include",
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    })
    .then(res => {
        if(res.ok){
            fetch("/course/" + data.courseid)
            .then(res => res.json())
            .then(course => {
                showCourse(course)
                resetForm()
            })
        }else{
            alert("Server error while adding course")
        }
    })
}

// ================= EDIT COURSE (FILL FORM) =================
let selectedRow = null

function editCourse(btn){
    selectedRow = btn.parentElement.parentElement

    document.getElementById("cid").value = selectedRow.cells[0].innerHTML
    document.getElementById("cname").value = selectedRow.cells[1].innerHTML

    let cid = selectedRow.cells[0].innerHTML

    let button = document.getElementById("button-add")
    button.innerHTML = "Update"
    button.setAttribute("onclick", `updateCourse("${cid}")`)
}

// ================= UPDATE COURSE =================
function updateCourse(oldCID){
    let data = getCourseFormData()

    fetch("/course/" + oldCID,{
        method: "PUT",
        credentials: "include",
        body: JSON.stringify(data),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    })
    .then(res =>{
        if(res.ok){
            selectedRow.cells[0].innerHTML = data.courseid
            selectedRow.cells[1].innerHTML = data.coursename

            let btn = document.getElementById("button-add")
            btn.innerHTML = "Add"
            btn.setAttribute("onclick", "addCourse()")

            selectedRow = null
            resetForm()
        }else{
            alert("Update failed")
        }
    })
}

// ================= DELETE COURSE =================
function deleteCourse(btn){
    if(confirm("Are you sure to Delete this course?")){
        let row = btn.parentElement.parentElement
        let cid = row.cells[0].innerHTML

        fetch("/course/" + cid,{
            method: "DELETE",
            credentials: "include"
        })
        .then(res => {
            if(res.ok){
                document.getElementById("myTable").deleteRow(row.rowIndex)
            }else{
                alert("Delete failed")
            }
        })
    }
}

// // ===============================
// // CONFIG
// // ===============================
// const BASE_URL = "http://localhost:8888/api";

// // ===============================
// // LOAD ALL COURSES ON PAGE LOAD
// // ===============================
// window.onload = function () {
//     loadCourses();
// };

// // ===============================
// // GET ALL COURSES
// // ===============================
// async function loadCourses() {
//     try {
//         const response = await fetch(BASE_URL + "/course/all", {
//             method: "GET",
//             credentials: "include"   // ⭐ send session cookie
//         });

//         if (!response.ok) {
//             throw new Error("Failed to fetch courses");
//         }

//         const result = await response.json();
//         displayCourses(result.data);

//     } catch (error) {
//         console.error("Error loading courses:", error);
//         alert("Unable to load courses. Please login again.");
//     }
// }

// // ===============================
// // DISPLAY COURSES IN TABLE
// // ===============================
// function displayCourses(courses) {
//     const tableBody = document.getElementById("courseTableBody");
//     tableBody.innerHTML = "";

//     courses.forEach(course => {
//         const row = `
//             <tr>
//                 <td>${course.courseid}</td>
//                 <td>${course.coursename}</td>
//                 <td>
//                     <button onclick="deleteCourse(${course.courseid})">
//                         Delete
//                     </button>
//                 </td>
//             </tr>
//         `;
//         tableBody.innerHTML += row;
//     });
// }

// // ===============================
// // ADD COURSE
// // ===============================
// async function addCourse() {
//     const cid = document.getElementById("cid").value;
//     const cname = document.getElementById("cname").value;

//     if (!cid || !cname) {
//         alert("Please fill all fields");
//         return;
//     }

//     const courseData = {
//         courseid: parseInt(cid),
//         coursename: cname
//     };

//     try {
//         const response = await fetch(BASE_URL + "/course/add", {
//             method: "POST",
//             credentials: "include", // ⭐ send cookie
//             headers: {
//                 "Content-Type": "application/json"
//             },
//             body: JSON.stringify(courseData)
//         });

//         const result = await response.json();

//         if (!response.ok) {
//             throw new Error(result.message || "Failed to add course");
//         }

//         alert("Course added successfully!");
//         document.getElementById("cid").value = "";
//         document.getElementById("cname").value = "";

//         loadCourses(); // refresh table

//     } catch (error) {
//         console.error("Add course error:", error);
//         alert("Error adding course");
//     }
// }

// // ===============================
// // DELETE COURSE
// // ===============================
// async function deleteCourse(id) {
//     if (!confirm("Are you sure you want to delete this course?")) return;

//     try {
//         const response = await fetch(BASE_URL + "/course/" + id, {
//             method: "DELETE",
//             credentials: "include"  // ⭐ send cookie
//         });

//         const result = await response.json();

//         if (!response.ok) {
//             throw new Error(result.message || "Delete failed");
//         }

//         alert("Course deleted successfully!");
//         loadCourses();

//     } catch (error) {
//         console.error("Delete error:", error);
//         alert("Error deleting course");
//     }
// }
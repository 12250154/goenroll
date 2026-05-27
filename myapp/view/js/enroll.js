//const API = "http://localhost:8080";

const API = "https://your-backend-name.onrender.com";

window.onload = function () {
    loadStudents();
    loadCourses();
    loadEnrollments();
};



// ⭐ Helper: handles both {data:[]} OR []
function getArray(resp){
    if(Array.isArray(resp)) return resp;
    if(Array.isArray(resp.data)) return resp.data;
    return [];
}



// ================= STUDENTS =================
async function loadStudents() {
    try {
        const res = await fetch(API + "/student/all");
        const json = await res.json();

        const students = getArray(json);  // ⭐ magic fix

        const select = document.getElementById("sid");
        select.innerHTML = "";

        students.forEach(s => {
            let option = document.createElement("option");
            option.value = s.stdid;
            option.text = s.stdid + " - " + s.fname + " " + s.lname;
            select.appendChild(option);
        });

    } catch (err) {
        console.error("Student load error:", err);
    }
}



// ================= COURSES =================
async function loadCourses() {
    try {
        const res = await fetch(API + "/course/all");
        const json = await res.json();

        const courses = getArray(json);  // ⭐ magic fix

        const select = document.getElementById("cid");
        select.innerHTML = "";

        courses.forEach(c => {
            let option = document.createElement("option");
            option.value = c.courseid;
            option.text = c.courseid + " - " + c.coursename;
            select.appendChild(option);
        });

    } catch (err) {
        console.error("Course load error:", err);
    }
}



// ================= ENROLLMENTS =================
async function loadEnrollments() {
    try {
        const res = await fetch(API + "/enroll/all");
        const json = await res.json();

        const enrollments = getArray(json); // ⭐ magic fix

        const table = document.getElementById("myTable");

        table.innerHTML = `
            <tr>
                <th>Student ID</th>
                <th>Course ID</th>
                <th>Date Enrolled</th>
            </tr>
        `;

        enrollments.forEach(e => {
            let row = table.insertRow();
            row.insertCell(0).innerHTML = e.stdid;
            row.insertCell(1).innerHTML = e.cid;
            row.insertCell(2).innerHTML = e.date;
        });

    } catch (err) {
        console.error("Enrollment load error:", err);
    }
}



// ================= ENROLL =================
function getEnrollData(){
    return {
        stdid: parseInt(document.getElementById("sid").value),
        cid: document.getElementById("cid").value
    }
}

async function addEnroll() {
    const data = getEnrollData();

    const res = await fetch(API + "/enroll", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(data)
    });

    if(res.status === 201){
        alert("Student Enrolled Successfully 🎉");
        loadEnrollments(); // refresh table
    }
    else if(res.status === 403){
        alert("Student already enrolled in this course");
    }
    else{
        alert("Enrollment failed");
    }
}

window.addEnroll = addEnroll;
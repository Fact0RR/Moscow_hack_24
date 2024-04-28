        const allStar = document.querySelectorAll('.rating .star')
        const ratingValue = document.querySelector('.rating input')
       
        allStar.forEach((item, idx)=> {
            item.addEventListener('click', function () {
                let click = 0
                ratingValue.value = idx + 1

                allStar.forEach(i=> {
                    i.classList.replace('bxs-star', 'bx-star')
                    i.classList.remove('active')
                })
                for(let i=0; i<allStar.length; i++) {
                    if(i <= idx) {
                        allStar[i].classList.replace('bx-star', 'bxs-star')
                        allStar[i].classList.add('active')
                    } else {
                        allStar[i].style.setProperty('--i', click)
                        click++
                    }
                }
            })
        })

function sendRewiew(){
    

    parent = document.getElementById("rating")

    sum = 0
    for(let i = 1;i<=5;i++){
        if (parent.children[i].className === "bx bxs-star star active"){
            sum += 1
        }
        
    }  
    console.log(sum)

    text = document.getElementById("rewiewText")
    
    textFromForm.value

    form = {
        grade: sum,
        text: text.value
    }

    alert("ок")
    fetch("/rewiew",
    {
        headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(form)
    })
    .then((response) => response.text())
    .then((data) => {
    console.log(data)
    console.log(typeof data)
  })

}
   
async function sendDataToServer(){
    textFromForm = document.getElementById("textForm")
    form = {
        text: textFromForm.value
    }

    fetch("/send",
    {
        headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(form)
    })
    .then((response) => response.json())
    .then((data) => {
    console.log(data)
    console.log(typeof data)
    addCourses(data) 
  })

}

function addCourses(data){
document.getElementById("num").textContent = data.length
// Получаем родительский элемент, куда будем добавлять новые div
var parentElement = document.getElementById('cs');


// Цикл для создания новых div элементов и их добавления в родительский элемент
for (var i = 0; i < data.length; i++) {
    var courseData = data[i];
    
    var newCourseDiv = document.createElement('div');
    newCourseDiv.classList.add('course');

    newCourseDiv.innerHTML = `
        <a href="${courseData.link}" class="link" onclick="window.open('${courseData.link}', '_blank')'></a>
        <div class="icon">
            <img src="../css/images/GBicon.jpeg" alt="" class="ava">
        </div>
        <div class="context">
            <div class="general">
                <div class="Cname">
                    <span class="type">Профильный курс</span>
                    <div class="prof">
                        <div class="name">${courseData.name}</div>
                        <img src="../css/images/more.svg" alt="" class="more">
                    </div>
                </div>
                <div class="tags">
                    <div class="tag">
                        <i class='bx bx-time-five'></i>
                        <span class="time">${courseData.duration}</span>
                    </div>
                </div>
            </div>
        </div>
    `;
    
    parentElement.appendChild(newCourseDiv);
}
}
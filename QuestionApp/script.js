let questions = [];
let score = 0;

function addQuestion() {
    const questionInput = document.getElementById('questionInput').value;
    if (questionInput) {
        const question = { text: questionInput, difficulty: Math.floor(Math.random() * 5) + 1 };
        questions.push(question);
        displayQuestions();
        document.getElementById('questionInput').value = '';
    }
}

function displayQuestions() {
    const questionList = document.getElementById('questionList');
    questionList.innerHTML = questions.map((q, index) => 
        `<p>${q.text} 
        <button onclick="editQuestion(${index})">Düzenle</button> 
        <button onclick="deleteQuestion(${index})">Sil</button></p>`).join('');
}

function deleteQuestion(index) {
    questions.splice(index, 1);
    displayQuestions();
}

function editQuestion(index) {
    const newQuestion = prompt("Yeni soruyu girin:", questions[index].text);
    if (newQuestion) {
        questions[index].text = newQuestion;
        displayQuestions();
    }
}

<script lang="ts">
  // import logo from './assets/images/logo-universal.png';
  import './app.css';
  import logo from '../../build/appicon.png';
  import { Quiz } from '../wailsjs/go/main/App.js';

  let resultText: string = 'Please drop your file here 👇';
  let name: string;
  let file: File;

  // function quiz(): void {
  //   // send file to backend
  //   Quiz(file).then((result) => (resultText = result));
  // }

  function quiz(): void {
    if (!file) {
      resultText = 'Please drop a file first';
      return;
    }

    const reader = new FileReader();
    reader.onload = () => {
      const fileContents = reader.result as string;
      Quiz(btoa(fileContents)).then((result) => (resultText = result));
    };
    reader.readAsBinaryString(file);
  }

  function handleDrop(event: DragEvent): void {
    event.preventDefault();
    file = event.dataTransfer.files[0];
    name = file.name;

    console.log(file);
    if (
      file.type !== 'application/vnd.openxmlformats-officedocument.wordprocessingml.document' &&
      file.type !== 'text/csv'
    ) {
      resultText = 'Seriously?! 🤨 Only DOCX or CSV files are allowed!';
    } else {
      resultText = `Exporting: ${name}`;
      quiz();
    }
  }
</script>

<main class="flex flex-col items-center">
  <img id="logo" class="w-40 p-4" alt="StoryQuiz Logo" src={logo} />
  <h1 class="text-xl font-black m-4">Storyline Quiz Exporter</h1>
  <div id="result" class="">{resultText}</div>
  <article
    id="drop"
    class="border-2 border-dotted p-6 m-4 rounded-lg"
    on:drop={handleDrop}
    on:dragover={(e) => e.preventDefault()}
    on:dragenter={(e) => {
      e.preventDefault();
      e.target.classList.add('bg-slate-600');
    }}
    on:dragleave={(e) => {
      e.preventDefault();
      e.target.classList.remove('bg-slate-600');
    }}
  >
    Drag & Drop a DOCX or CSV file
  </article>
</main>

<style>
  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style>

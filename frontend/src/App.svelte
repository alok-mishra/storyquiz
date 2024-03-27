<script lang="ts">
    // import logo from './assets/images/logo-universal.png';

    import storyquiz from './assets/images/storyquiz.svg';
    import storyline from './assets/images/storyline.svg';
    import csod from './assets/images/csod.svg';
    import excel from './assets/images/excel.svg';
    import word from './assets/images/word.svg';
    import csv from './assets/images/csv.svg';
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
    <!-- <img id="logo" class="w-40 p-4" alt="StoryQuiz Logo" src={logo} /> -->
    <h1 class="text-xl font-black m-4">Quiz Exporter</h1>
    <section class="relative">
        <img class="w-60 p-4" alt="StoryQuiz Logo" src={storyquiz} />
        <img class="w-40 p-4 absolute -bottom-0 -left-20" alt="Storyline Logo" src={storyline} />
        <img class="w-40 p-4 absolute -bottom-0 -right-20" alt="CSOD Logo" src={csod} />
    </section>
    <div id="result" class="">{resultText}</div>
    <section class="dragdrop flex">
        <article
            id="drop"
            class="border-2 border-dotted p-6 m-4 rounded-lg bg-opacity-60 max-w-60"
            on:drop={handleDrop}
            on:dragover={(e) => e.preventDefault()}
            on:dragenter={(e) => {
                e.preventDefault();
                e.target.classList.add('bg-[#c560b3]');
            }}
            on:dragleave={(e) => {
                e.preventDefault();
                e.target.classList.remove('bg-[#c560b3]');
            }}
        >
            Drag & Drop a Word (.docx) or CSV file here
        </article>

        <article
            id="drop"
            class="border-2 border-dotted p-6 m-4 rounded-lg bg-opacity-60 max-w-60"
            on:drop={handleDrop}
            on:dragover={(e) => e.preventDefault()}
            on:dragenter={(e) => {
                e.preventDefault();
                e.target.classList.add('bg-[#fa4616]');
            }}
            on:dragleave={(e) => {
                e.preventDefault();
                e.target.classList.remove('bg-[#fa4616]');
            }}
        >
            Drag & Drop an Excel (.xlsx) file here
        </article>
    </section>
</main>

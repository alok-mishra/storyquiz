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

    import DragDrop from './components/DragDrop.svelte';

    let message: string = '👇 Please drop your file here 👇';
    let name: string;
    let file: File;

    function onDrop(file: File, resultText: string): void {
        message = resultText;

        if (file) {
            console.log('reading file...', file);
            name = file.name;
            const reader = new FileReader();
            reader.onload = () => {
                const fileContents = reader.result as string;
                Quiz(btoa(fileContents)).then((result) => {
                    message = result;
                });
            };
            reader.readAsBinaryString(file);
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
    <div id="message" class="">{message}</div>
    <section class="dragdrop flex">
        <DragDrop dropColor={'#c560b3'} fileTypes={['docx', 'xlsm']} {onDrop}>
            Drop a Word or Excel file here for Storyline Quiz
        </DragDrop>
        <DragDrop dropColor={'#fa4616'} fileTypes={['xls', 'xlsx', 'xlsm']} {onDrop}>
            Drop a Excel file here for CSOD Quiz
        </DragDrop>
    </section>
</main>

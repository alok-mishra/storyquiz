<div style="display: flex; justify-content: center; align-items: center;">
  <img src="https://raw.githubusercontent.com/alok-mishra/storyquiz/master/frontend/src/assets/images/storyquiz.svg" width="100" height="100" />
  <span style="font-size: 4em; font-weight: bold; padding: 20px;">StoryQuiz</span>
</div>

## Overview

The StoryQuiz application is a tool designed to streamline the process of creating quizzes for Articulate Storyline and Cornerstone Learning Management Systems. Users can drag and drop Word or Excel files into the interface, and the application will automatically process these files to extract questions and generate the corresponding import files.

## Usage Instructions

1. **Drag-and-Drop**: Drag your Word or Excel file into the designated area of the interface. You can also drop multiple files simultaneously, and the application will process them concurrently.
2. **Input File**: You can choose from course scripts with embedded test questions, a separate test script, or the legacy Excel test format.
3. **File Processing**: The app will automatically process the file(s) and generate the necessary quiz import files.
4. **Output**: Check the output files for quiz import into Storyline or CSOD. The output files will be saved in the same directory where the app executable resides.

### Extracting Course Codes

-   Ensure your Word or Excel files contain the course code in the standard location.
-   The app will automatically detect and extract these codes for use in the output files.

### Viewing Version Information

-   Hover over the version displayed in the header to see the latest release notes.

## Output File Naming Convention

-   The output files will have the same name as the original script.
-   For Storyline, the output will be a txt file with (Storyline) appended to the end of the filename.
-   For Cornerstone, the output will be in Excel format with (LMS) appended to the end of the filename.

## Error Handling

The app includes several error handling mechanisms to ensure a smooth user experience:

-   **Error Displays**: Errors are displayed in an easy format to help you quickly identify and resolve issues.
-   **Progress Meter**: While processing files, a progress meter is displayed to keep you informed about the status.

## Importing Quizzes into Storyline

1. **Import Questions**: When importing the generated quiz file into Storyline, import the questions to a new scene.
2. **Apply Layout**: Select all the slides in that scene and apply the "Questions with Headers" layout.
3. **Set Feedback**: Set each slide to "Feedback: None" under the Question Tools Design tab.
4. **Move Slides**: Select all of the slides and move them after slide 1.2.
5. **Delete Empty Scene**: Delete the now-empty scene.
6. **Review for Errors**: Review the slides for any errors and correct them as necessary.

## Keeping Up-to-Date

To ensure compatibility and benefit from the latest features and improvements, it is recommended that users update to the latest version of the app every few months as standards and requirements change.

## Contact and Support

For further assistance or to report issues, please use the information provided in the app.

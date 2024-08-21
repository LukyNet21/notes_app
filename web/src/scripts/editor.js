import EditorJS from "@editorjs/editorjs";
import Header from "@editorjs/header";
import SimpleImage from "@editorjs/simple-image";
import Checklist from '@editorjs/checklist'
import List from "@editorjs/list";
import Table from '@editorjs/table'
import Embed from '@editorjs/embed';
import Underline from '@editorjs/underline';
import CodeTool from '@editorjs/code';
import InlineCode from '@editorjs/inline-code';
import Warning from '@editorjs/warning';
import Marker from '@editorjs/marker';
import Quote from '@editorjs/quote';

const editor = new EditorJS({
    theme: "dark",
    tools: {
        header: Header,
        image: SimpleImage,
        embed: Embed,
        table: Table,
        underline: Underline,
        code: CodeTool,
        quote: Quote,
        checklist: {
            class: Checklist,
            inlineToolbar: true,
        },
        list: {
            class: List,
            inlineToolbar: true,
            config: {
                defaultStyle: 'unordered'
            }
        },
        inlineCode: {
            class: InlineCode,
            shortcut: 'CMD+SHIFT+M',
        },
        warning: {
            class: Warning,
            inlineToolbar: true,
            shortcut: 'CMD+SHIFT+W',
            config: {
                titlePlaceholder: 'Title',
                messagePlaceholder: 'Message',
            },
        },
        marker: {
            class: Marker,
            shortcut: 'CMD+SHIFT+M',
        }
    },
})

const queryString = window.location.search;
const urlParams = new URLSearchParams(queryString);
const id = urlParams.get('id')
const noteNameEditInput = document.getElementById("noteNameEditInput")
const noteNameField = document.getElementById("projectName")

let noteName

editor.isReady.then(() => {
    fetch(`http://localhost:8080/getNoteByID/${id}`, {
        method: "GET",
        credentials: "include",
    })
        .then((response) => {
            if (!response.ok) {
                window.location.replace("/dashboard");
                throw new Error("Network response was not ok");
            }
            return response.json();
        })
        .then((data) => {
            console.log("Success:", data);
            noteName = data.note_name
            noteNameField.innerText = noteName
            noteNameEditInput.value = noteName
            editor.blocks.render(JSON.parse(data.note_content))
        })
        .catch((error) => {
            console.error("Error:", error);
        });
})
    .catch((reason) => {
        console.log(`Editor.js initialization failed because of ${reason}`)
    });

document.getElementById("saveButton").addEventListener("click", () => {
    editor.save().then((outputData) => {
        console.log('Article data: ', JSON.stringify(outputData))
        fetch(`http://localhost:8080/updateNote/${id}`, {
            method: "PUT",
            credentials: "include",
            body: JSON.stringify({
                "note_name": noteName,
                "note_content": JSON.stringify(outputData)
            })
        })
    })
        .catch((error) => {
            console.log('Saving failed: ', error)
        });
})

document.getElementById("noteNameEditBtn").addEventListener("click", e => {
    document.getElementById("projectName").classList.toggle("hidden")
    noteNameEditInput.classList.toggle("hidden")
})

noteNameEditInput.addEventListener("change", e => {
    noteName = e.target.value
    noteNameField.innerText = noteName
})

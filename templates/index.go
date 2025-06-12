package templates

import "html/template"

var Temp = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>House To-Do List</title>
        <style>
body {
    font-family: Arial, sans-serif;
    padding: 20px;
    background-color: #f5f5f5;
}

    h1 {
        text-align: center;
    }

    .item-list {
        max-width: 1000px;
        width: 300px;
        padding: 0;
    }

    .item {
        position: relative;
        background: white;
        padding: 1rem;
        margin-bottom: 1rem;
        border-radius: 10px;
        box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
    }

    .item-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .item span {
        display: block;
        margin: 0.25rem 0;
    }

    .item span strong {
        color: #555;
        width: 80px;
        display: inline-block;
    }
    .form-container {
        max-width: 500px;
        width: 300px;
        background: white;
        padding: 2rem;
        border-radius: 12px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    }

    .form-container h2 {
        text-align: center;
        margin-bottom: 1.5rem;
        color: #333;
    }

    .form-group {
        margin-bottom: 1rem;
    }

    .form-group label {
        display: block;
        font-weight: bold;
        margin-bottom: 0.5rem;
        color: #555;
    }

    .form-group input {
        width: 100%;
        padding: 0.6rem;
        border: 1px solid #ccc;
        border-radius: 8px;
        font-size: 1rem;
    }

    .form-container button {
        width: 100%;
        padding: 0.75rem;
        background: #007bff;
        color: white;
        font-size: 1rem;
        border: none;
        border-radius: 8px;
        cursor: pointer;
        transition: background 0.3s ease;
    }

    .form-container button:hover {
        background: #0056b3;
    }
    .content-container {
        display: flex;
        justify-content: center;
        align-items: flex-start;
        gap: 2rem;
        max-width: 1200px;
        margin: 2rem auto;
        padding: 0 1rem;
    }
    .item-btn {
        background: #ff4d4d;
        border: none;
        color: white;
        font-size: 0.7rem;
        padding: 0.2rem 0.4rem;
        border-radius: 6px;
        cursor: pointer;
        transition: background 0.2s ease;
    }

    .item-btn:hover {
        background: #cc0000;
    }
        </style>
    </head>
    <body>

        <h1>House Grocery List</h1>

        <div class="content-container">
            <div class="form-container">
                <h2>Add Grocery Item</h2>
                <form action="/form" method="POST">
                    <div class="form-group">
                        <label for="item">Item Name:</label>
                        <input type="text" name="item" id="item" required />
                    </div>

                    <div class="form-group">
                        <label for="user">Person Responsible:</label>
                        <input type="text" name="user" id="user" required />
                    </div>

                    <div class="form-group">
                        <label for="location">Location:</label>
                        <input type="text" name="location" id="location" required />
                    </div>

                    <button type="submit">Add Item</button>
                </form>
            </div>

            <br>
            <div class="item-list" id="itemList">
            </div>
        </div>
        <script>
            async function loadItems() {
                try {
                    const response = await fetch('/items'); 
                    const items = await response.json();

                    const itemList = document.getElementById('itemList');
                    itemList.innerHTML = '';

                    items.forEach(item => {
                        const div = document.createElement('div');
                        div.className = 'item';
                        div.innerHTML = '<div class="item-header">' +
                    '<span><strong>What:</strong> ' + item.name + '</span>' +
                    '<button class="item-btn" onclick="deleteItem(' + item.id + ')">✖</button>' +
                    '</div>' +
                    '<span><strong>Who:</strong> ' + item.user + '</span>' +
                    '<span><strong>Where:</strong> ' + item.location + '</span>';
                        itemList.appendChild(div);
                    });
                } catch (error) {
                    console.error('Failed to load items:', error);
                }
            }

            async function deleteItem(id) {

                try {
                    const res = await fetch('/delete/' + id, {
                        method: 'DELETE'
                    });

                    if (!res.ok) {
                        const msg = await res.json();
                        alert("Failed to delete: " + (msg.error || "Unknown error"));
                        return;
                    }

                    loadItems(); // Refresh list
                } catch (err) {
                    console.error("Error deleting item:", err);
                    alert("Error deleting item");
                }
            }


            loadItems();
        </script>
    </body>

</html>
`))

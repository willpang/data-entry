function sortTable(n, data_type) {
  var table,
    rows,
    switching,
    i,
    x,
    y,
    shouldSwitch,
    dir,
    switchcount = 0;
  table = document.getElementById("dataTable");
  switching = true;
  // Set the sorting direction to ascending:
  dir = "asc";
  /* Make a loop that will continue until
  no switching has been done: */
  while (switching) {
    // Start by saying: no switching is done:
    switching = false;
    rows = table.rows;
    /* Loop through all table rows (except the
    first, which contains table headers): */
    for (i = 1; i < rows.length - 1; i++) {
      // Start by saying there should be no switching:
      shouldSwitch = false;
      /* Get the two elements you want to compare,
      one from current row and one from the next: */
      x = rows[i].getElementsByTagName("TD")[n];
      y = rows[i + 1].getElementsByTagName("TD")[n];
      /* Check if the two rows should switch place,
      based on the direction, asc or desc: */
      if (dir == "asc") {
        if (data_type == "int") {
          if (parseInt(x.innerHTML) > parseInt(y.innerHTML)) {
            // If so, mark as a switch and break the loop:
            shouldSwitch = true;
            break;
          }
        } else if (data_type == "date") {
          if (Date.parse(x.innerHTML) > Date.parse(y.innerHTML)) {
            // If so, mark as a switch and break the loop:
            shouldSwitch = true;
            break;
          }
        } else if (data_type == "float") {
          if (parseFloat(x.innerHTML) > parseFloat(y.innerHTML)) {
            // If so, mark as a switch and break the loop:
            shouldSwitch = true;
            break;
          }
        } else if (data_type == "string") {
          if (x.innerHTML.toLowerCase() > y.innerHTML.toLowerCase()) {
            // If so, mark as a switch and break the loop:
            shouldSwitch = true;
            break;
          }
        }
      } else if (dir == "desc") {
        if (data_type == "int") {
          if (parseInt(x.innerHTML) < parseInt(y.innerHTML)) {
            // If so, mark as a switch and break the loop:
            shouldSwitch = true;
            break;
          }
        } else if (data_type == "date") {
          if (Date.parse(x.innerHTML) < Date.parse(y.innerHTML)) {
            // If so, mark as a switch and break the loop:
            shouldSwitch = true;
            break;
          }
        } else if (data_type == "float") {
          if (parseFloat(x.innerHTML) < parseFloat(y.innerHTML)) {
            // If so, mark as a switch and break the loop:
            shouldSwitch = true;
            break;
          }
        } else if (data_type == "string") {
          if (x.innerHTML.toLowerCase() < y.innerHTML.toLowerCase()) {
            // If so, mark as a switch and break the loop:
            shouldSwitch = true;
            break;
          }
        }
      }
    }
    if (shouldSwitch) {
      /* If a switch has been marked, make the switch
      and mark that a switch has been done: */
      rows[i].parentNode.insertBefore(rows[i + 1], rows[i]);
      switching = true;
      // Each time a switch is done, increase this count by 1:
      switchcount++;
    } else {
      /* If no switching has been done AND the direction is "asc",
      set the direction to "desc" and run the while loop again. */
      if (switchcount == 0 && dir == "asc") {
        dir = "desc";
        switching = true;
      }
    }
  }
}

function searchTable() {
  // Declare variables
  var input, filter, table, tr, td, i, txtValue;
  input = document.getElementById("tableSearchInput");
  filter = input.value.toUpperCase();
  table = document.getElementById("dataTable");
  tr = table.getElementsByTagName("tr");

  // Loop through all table rows, and hide those who don't match the search query
  for (i = 0; i < tr.length; i++) {
    td = tr[i].getElementsByTagName("td");
    found = false;
    for (j = 0; j < td.length; j++) {
      val = td[j];
      if (val && !found) {
        txtValue = val.textContent || val.innerText;
        if (txtValue.toUpperCase().indexOf(filter) > -1) {
          tr[i].style.display = "";
          found = true;
        } else {
          tr[i].style.display = "none";
        }
      }
    }
  }
}

function submitCreateUserForm(event) {
  event.preventDefault();
  const form = document.getElementById("create-user-form");
  const formData = new FormData(form);
  const userData = Object.fromEntries(formData.entries());
  const options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(userData),
  };
  fetch("/user", options)
    .then((response) => {
      if (response.ok) {
        // Display success message
        alert("Pelanggan telah sukses didaftarkan!");
        // Redirect to user list page
        window.location.href = "/users";
      } else {
        throw new Error("Gagal mendaftarkan pelanggan");
      }
    })
    .catch((error) => console.error(error));
}

function deleteUser(id, name) {
  var r = confirm("Apakah anda yakin ingin menghapus pelanggan ini?");
  if (r == true) {
    fetch("/user/" + id, {
      method: "DELETE",
    })
      .then((response) => {
        if (response.ok) {
          alert(
            "Sukses menghapus pelanggan " + name + " dengan id " + id + "..."
          );
          location.reload(); // Reload the page after the user is deleted
        } else {
          alert("Gagal menghapus pelanggan " + name + " dengan id " + id + "!");
        }
      })
      .catch((error) => {
        alert("Gagal menghapus pelanggan " + name + ": " + error.message);
      });

    return true;
  } else {
    return false;
  }
}

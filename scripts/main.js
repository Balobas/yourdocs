async function searchDocuments(categories, names, docNumbers) {
	if (!categories.isArray()) {
		return false
	}
	if (!names.isArray()) {

	}
	let response = await fetch("localhost:8089", {
		method: "POST",
		headers: "content-type: application/json",
		body: {
			categories: categories,
			names: names,
			docNumbers: docNumbers,
		}
	});
	let docs = await response.json();
}

async function getDocument(uid)
{
	let response = await fetch ("localhost:8089?uid=" + uid, {
		method: "GET",
	});
	let doc = await response.json();
}

async function getCategories(page, catOnPage)
{

}
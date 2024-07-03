# artigo

A simple bespoke REST API built using Go that leverages [Auth0](https://auth0.com) for authentication & authorization.

## Stack

* Authz/Authn: Auth0
* Database: MongoDB

## Endpoints

<table>
<thead>
<tr>
<th>Verb</th>
<th>URI</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>GET</td>
<td><code>/articles</code></td>
<td>Fetches all articles from the database</td>
</tr>
 <tr>
<tr>
<td>POST</td>
<td><code>/articles</code></td>
<td>Creates a new article record in the database</td>
</tr>
<td>GET</td>
<td><code>/articles/{id}</code></td>
<td>Retrieves an article with the given ID</td>
</tr>
<tr>
<tr>
<td>PUT</td>
<td><code>/articles/{id}</code></td>
<td>Updates the given article</td>
</tr>
<td>DELETE</td>
<td><code>/articles/{id}</code></td>
<td>Deletes the given article</td>
</tr>
</tbody>
</table>

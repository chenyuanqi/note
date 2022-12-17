<?php

use App\Http\Resources\PostResource;
use App\Models\Post;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use Spatie\QueryBuilder\QueryBuilder;

use function Termwind\ask;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

Route::get('posts/{id}', function (Request $request, int $id) {
    $post =  QueryBuilder::for(Post::where('id', $id))
        ->allowedFields(['id', 'title', 'content', 'slug', 'views', 'created_at', 'authors.id', 'authors.name'])
        ->allowedIncludes(['author', 'comments'])
        ->first();
    return new PostResource($post);
})->name('posts.show');

Route::get('posts', function (Request $request) {
    $posts = QueryBuilder::for(Post::class)
        ->allowedFields(['id', 'title', 'content', 'views', 'created_at', 'authors.id', 'authors.name'])
        ->allowedFilters(['title'])
        ->defaultSort('-id')
        ->allowedSorts(['views', 'created_at'])
        ->allowedIncludes('author')
        ->paginate(10);
    return PostResource::collection($posts);
});

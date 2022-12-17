<?php
namespace App\Http\Controllers;

use App\Actions\PaydayAction;
use Illuminate\Http\Request;

class PaycheckController extends Controller
{
    public function __construct(private readonly PaydayAction $payday)
    {
    }

    public function store(Request $request) 
    {
       $this->payday->execute();
       return response()->noContent();
    }
}

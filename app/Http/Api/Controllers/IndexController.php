<?php

namespace App\Http\Api\Controllers;

class IndexController extends BaseController
{
    public function index()
    {
        return [
            'version' => 1,
            'status' => 'ok',
        ];
    }
}

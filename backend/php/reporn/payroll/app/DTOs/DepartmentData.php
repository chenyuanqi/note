<?php

namespace App\DTOs;

class DepartmentData
{
    public function __construct(
        public readonly string $name,
        public readonly ?string $description
    ) {}
}
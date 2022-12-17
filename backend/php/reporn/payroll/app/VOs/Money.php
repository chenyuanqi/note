<?php
namespace App\VOs;

class Money
{
    public function __construct(private readonly int $valueInCents)
    {
        
    }

    public static function from(int $valueCents): self
    {
        return new static($valueCents);
    }

    public function toDollars(): string 
    {
        return '$' . number_format($this->valueInCents / 100, 2);
    }

    public function toCents(): int
    {
        return $this->valueInCents;
    }
}
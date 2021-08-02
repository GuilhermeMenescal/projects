from django.shortcuts import render
from .models import Publication

def index(request):
    return render(request, 'index.html')

def publications(request):
    publis = Publication.objects.all()
    context={'publis' : publis}
    return render(request, 'publications.html', context)

# This is a sample Python script.

# Press Shift+F10 to execute it or replace it with your code.
# Press Double Shift to search everywhere for classes, files, tool windows, actions, and settings.
from tts import tts
from stt import stt
def start():
    valor = ''
    print('Digite TTS para usar o texto para voz, ou STT para usar voz para texto: ')
    print('Se desejar encerrar, digite Fim')
    valor = input('>> ')

    if( valor == 'TTS'):
            tts()
            start()
    elif(valor == 'STT'):
            stt()
            start()
    elif (valor == 'Fim'):
            exit()
    else:
        print('Por favor digite um dos valores pedidos!')
        start()
if __name__ == '__main__':
    start()
# Press the green button in the gutter to run the script.

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
